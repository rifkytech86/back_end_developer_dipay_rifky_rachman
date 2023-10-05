package controller

import (
	"context"
	"github.com/dipay/api"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/dipay/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type userAdminController struct {
	UserAdminUseCase usecase.IUseCaseUserAdmin
	ContextTimeOut   int
	Validator        validations.IValidator
}

type IUserAdminController interface {
	Hello(ctx echo.Context, param api.HelloParams) error
	Login(ctx echo.Context) error
}

func NewUserAdminController(userAdminUseCase usecase.IUseCaseUserAdmin, contextTimeOut int, validator validations.IValidator) IUserAdminController {
	return &userAdminController{
		UserAdminUseCase: userAdminUseCase,
		ContextTimeOut:   contextTimeOut,
		Validator:        validator,
	}
}

func (u *userAdminController) Hello(ctx echo.Context, param api.HelloParams) error {
	return ctx.JSON(http.StatusOK, "sadfasd")
}

func (u *userAdminController) Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(u.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(api.LoginJSONBody)
	if err := c.Bind(req); err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	validatorErrors := u.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, validations.GetCustomMessage(err.Error, err.Field))
		}
		return commons.ErrorResponses(c, http.StatusBadRequest, internal.ErrorInvalidRequest.String(), errorMessages)
	}

	token, err := u.UserAdminUseCase.Login(ctx, req)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.LoginResponse{
		Code:   commons.String(strconv.Itoa(http.StatusCreated)),
		Status: commons.Int(http.StatusCreated),
		Data: &struct {
			Token *string `json:"token,omitempty"`
		}{
			Token: commons.String(token),
		},
		Message: commons.String(internal.SuccessMessage),
	}

	return c.JSON(http.StatusOK, response)
}
