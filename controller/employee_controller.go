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

type employeeController struct {
	EmployeeUseCase usecase.IEmployeeUseCase
	ContextTimeOut  int
	Validator       validations.IValidator
}

//go:generate mockery --name IEmployeeController
type IEmployeeController interface {
	AddEmployee(c echo.Context, companyId string) error
	GetEmployeeByID(c echo.Context, employeeID string) error
	GetEmployeeByCompanyID(ctx echo.Context, companyID string) error
	UpdateEmployeeData(ctx echo.Context, companyId string, employeeId string) error
	DeleteEmployeeByID(ctx echo.Context, employeeId string) error
}

func NewEmployeeController(employeeUseCase usecase.IEmployeeUseCase, contextTimeOut int, validator validations.IValidator) IEmployeeController {
	return &employeeController{
		EmployeeUseCase: employeeUseCase,
		ContextTimeOut:  contextTimeOut,
		Validator:       validator,
	}
}

func (e *employeeController) AddEmployee(c echo.Context, companyId string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(e.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(api.AddEmployeeJSONBody)
	if err := c.Bind(req); err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	validatorErrors := e.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, validations.GetCustomMessage(err.Error, err.Field))
		}
		return commons.ErrorResponses(c, http.StatusBadRequest, internal.ErrorInvalidRequest.String(), errorMessages)
	}

	employeeID, companyID, err := e.EmployeeUseCase.AddEmployee(ctx, companyId, req)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.AddEmployeeResponse{
		Code:   commons.String(strconv.Itoa(http.StatusCreated)),
		Status: commons.Int(http.StatusCreated),
		Data: &struct {
			CompanyId *string `json:"company_id,omitempty"`
			Id        *string `json:"id,omitempty"`
		}{
			CompanyId: &companyID,
			Id:        &employeeID,
		},
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusCreated, response)
}

func (e *employeeController) GetEmployeeByID(c echo.Context, employeeID string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(e.ContextTimeOut)*time.Second)
	defer cancel()

	if employeeID == "" {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}
	employee, err := e.EmployeeUseCase.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.GetEmployeeByIDResponse{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK),
		Data: &struct {
			Id          *string `json:"id,omitempty"`
			Jobtitle    *string `json:"jobtitle,omitempty"`
			Name        *string `json:"name,omitempty"`
			PhoneNumber *string `json:"phone_number,omitempty"`
		}{
			Id:          commons.String(employee.ID.Hex()),
			Name:        commons.String(employee.Name),
			Jobtitle:    commons.String(employee.JobTitle),
			PhoneNumber: commons.String(employee.PhoneNumber),
		},
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusOK, response)
}

func (e *employeeController) GetEmployeeByCompanyID(c echo.Context, companyID string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(e.ContextTimeOut)*time.Second)
	defer cancel()

	if companyID == "" {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	listEmployee, err := e.EmployeeUseCase.GetEmployeeByCompanyID(ctx, companyID)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}
	var convertRows []struct {
		Id          *string `json:"id,omitempty"`
		Jobtitle    *string `json:"jobtitle,omitempty"`
		Name        *string `json:"name,omitempty"`
		PhoneNumber *string `json:"phone_number,omitempty"`
	}
	for _, employee := range listEmployee {
		id := employee.ID.Hex()
		convertRows = append(convertRows, struct {
			Id          *string `json:"id,omitempty"`
			Jobtitle    *string `json:"jobtitle,omitempty"`
			Name        *string `json:"name,omitempty"`
			PhoneNumber *string `json:"phone_number,omitempty"`
		}{
			Id:          &id,
			Jobtitle:    &employee.JobTitle,
			Name:        &employee.Name,
			PhoneNumber: &employee.PhoneNumber,
		})
	}

	response := api.GetEmployeeByCompanyIDResponse{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK), // Assuming you have a helper function String to create pointers to strings
		Data: &struct {
			Employees *[]struct {
				Id          *string `json:"id,omitempty"`
				Jobtitle    *string `json:"jobtitle,omitempty"`
				Name        *string `json:"name,omitempty"`
				PhoneNumber *string `json:"phone_number,omitempty"`
			} `json:"employees,omitempty"`
		}{
			Employees: &convertRows,
		},
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusOK, response)
}

func (e *employeeController) UpdateEmployeeData(c echo.Context, companyID string, employeeId string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(e.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(api.UpdateEmployeeDataJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	validatorErrors := e.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, validations.GetCustomMessage(err.Error, err.Field))
		}
		return commons.ErrorResponses(c, http.StatusBadRequest, internal.ErrorInvalidRequest.String(), errorMessages)
	}

	resEmployeeID, resCompanyID, err := e.EmployeeUseCase.UpdateEmployeeData(ctx, companyID, employeeId, req)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.UpdateEmployeeDataResponse{
		Code:   commons.String(strconv.Itoa(http.StatusCreated)),
		Status: commons.Int(http.StatusCreated),
		Data: &struct {
			CompanyId *string `json:"company_id,omitempty"`
			Id        *string `json:"id,omitempty"`
		}{
			CompanyId: &resEmployeeID,
			Id:        &resCompanyID,
		},
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusCreated, response)
}

func (e *employeeController) DeleteEmployeeByID(c echo.Context, employeeId string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(e.ContextTimeOut)*time.Second)
	defer cancel()

	if employeeId == "" {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.ErrorInvalidRequest.GetCode(), internal.ErrorInvalidRequest.String())
	}

	err := e.EmployeeUseCase.DeleteEmployeeByID(ctx, employeeId)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}
