package controller

import (
	"context"
	"github.com/dipay/api"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type duplicateZeroController struct {
	ContextTimeOut int
	Validator      validations.IValidator
}

type IDuplicateZeroController interface {
	DuplicateZero(ctx echo.Context) error
}

func NewDuplicateZeroController(contextTimeOut int, validator validations.IValidator) IDuplicateZeroController {
	return &duplicateZeroController{
		ContextTimeOut: contextTimeOut,
		Validator:      validator,
	}
}

func (d *duplicateZeroController) DuplicateZero(c echo.Context) error {
	_, cancel := context.WithTimeout(c.Request().Context(), time.Duration(d.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(api.DuplicateZeroJSONBody)
	if err := c.Bind(req); err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	validatorErrors := d.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, validations.GetCustomMessage(err.Error, err.Field))
		}
		return commons.ErrorResponses(c, http.StatusBadRequest, internal.ErrorInvalidRequest.String(), errorMessages)
	}
	res := internal.CheckDuplicateZero(req.N)
	response := api.DuplicateZeroResponse{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK),
		Data: (*struct {
			Result *[]int32 `json:"result,omitempty"`
		})(&struct {
			Result *[]int32 `json:"result,omitempty"`
		}{
			Result: &res,
		}),
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusOK, response)
}
