package controller

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/dipay/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

type userAdminController struct {
	UserAdminUseCase usecase.IUseCaseUserAdmin
	ContextTimeOut   int
	Validator        validations.IValidator
}

//go:generate mockery --name IUserAdminController
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
	return ctx.JSON(http.StatusOK, "Hello Di-pay")
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			res, err := u.CreateUser(c, req)
			if err != nil {
				return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
			}
			return c.JSON(http.StatusCreated, res)
		}
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.LoginResponse{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK),
		Data: &struct {
			Token *string `json:"token,omitempty"`
		}{
			Token: commons.String(token),
		},
		Message: commons.String(internal.SuccessMessage),
	}

	return c.JSON(http.StatusOK, response)
}

func (u *userAdminController) CreateUser(c echo.Context, req *api.LoginJSONBody) (*api.LoginResponse, error) {
	token, err := u.UserAdminUseCase.Register(c.Request().Context(), req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New(internal.ErrorInternalServer.String())
		}
		return nil, err
	}
	response := &api.LoginResponse{
		Code:   commons.String(strconv.Itoa(http.StatusCreated)),
		Status: commons.Int(http.StatusCreated),
		Data: &struct {
			Token *string `json:"token,omitempty"`
		}{
			Token: commons.String(token),
		},
		Message: commons.String(internal.SuccessMessage),
	}

	return response, nil

}
