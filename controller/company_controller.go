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

type companyController struct {
	CompanyUseCase usecase.ICompanyUseCase
	ContextTimeOut int
	Validator      validations.IValidator
}

//go:generate mockery --name ICompanyController
type ICompanyController interface {
	AddCompany(ctx echo.Context) error
	GetCompany(ctx echo.Context) error
	UpdateCompanyStatusActive(ctx echo.Context, id string) error
}

func NewCompanyController(companyUseCase usecase.ICompanyUseCase, contextTimeOut int, validator validations.IValidator) ICompanyController {
	return &companyController{
		CompanyUseCase: companyUseCase,
		ContextTimeOut: contextTimeOut,
		Validator:      validator,
	}
}

func (u *companyController) AddCompany(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(u.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(api.AddCompanyJSONBody)
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

	lastInsertID, err := u.CompanyUseCase.AddCompany(ctx, req)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}
	response := api.CompaniesResponse{
		Code:   commons.String(strconv.Itoa(http.StatusCreated)),
		Status: commons.Int(http.StatusCreated),
		Data: (*struct {
			Id *string `json:"id,omitempty"`
		})(&struct {
			Id *string `json:"Id,omitempty"`
		}{
			Id: commons.String(lastInsertID),
		}),
		Message: commons.String(internal.SuccessMessage),
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *companyController) GetCompany(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(u.ContextTimeOut)*time.Second)
	defer cancel()
	listCompany, err := u.CompanyUseCase.GetCompany(ctx)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}
	var convertedRows []struct {
		Address         *string `json:"address,omitempty"`
		CompanyName     *string `json:"company_name,omitempty"`
		Id              *string `json:"id,omitempty"`
		IsActive        *bool   `json:"is_active,omitempty"`
		TelephoneNumber *string `json:"telephone_number,omitempty"`
	}

	for _, company := range listCompany {
		id := company.ID.Hex()
		convertedRows = append(convertedRows, struct {
			Address         *string `json:"address,omitempty"`
			CompanyName     *string `json:"company_name,omitempty"`
			Id              *string `json:"id,omitempty"`
			IsActive        *bool   `json:"is_active,omitempty"`
			TelephoneNumber *string `json:"telephone_number,omitempty"`
		}{
			Address:         &company.Address,
			CompanyName:     &company.CompanyName,
			Id:              &id,
			IsActive:        &company.IsActive,
			TelephoneNumber: &company.TelephoneNumber,
		})
	}

	response := api.GetCompanyResponse{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK),
		Data: &struct {
			Count *int `json:"count,omitempty"`
			Rows  *[]struct {
				Address         *string `json:"address,omitempty"`
				CompanyName     *string `json:"company_name,omitempty"`
				Id              *string `json:"id,omitempty"`
				IsActive        *bool   `json:"is_active,omitempty"`
				TelephoneNumber *string `json:"telephone_number,omitempty"`
			} `json:"rows,omitempty"`
		}{
			Count: commons.Int(len(listCompany)),
			Rows:  &convertedRows,
		},
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusOK, response)
}

func (u *companyController) UpdateCompanyStatusActive(c echo.Context, id string) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(u.ContextTimeOut)*time.Second)
	defer cancel()
	idAffected, isActive, err := u.CompanyUseCase.UpdateCompanyStatusActive(ctx, id)
	if err != nil {
		return commons.ErrorResponse(c, http.StatusBadRequest, internal.GetCodeByString(err.Error()), err.Error())
	}

	response := api.UpdateSetActiveCompany{
		Code:   commons.String(strconv.Itoa(http.StatusOK)),
		Status: commons.Int(http.StatusOK),
		Data: (*struct {
			Id       *string `json:"id,omitempty"`
			IsActive *bool   `json:"is_active,omitempty"`
		})(&struct {
			Id       *string `json:"Id,omitempty"`
			IsActive *bool   `json:"is_active,omitempty"`
		}{
			Id:       commons.String(idAffected),
			IsActive: commons.Boolean(isActive),
		}),
		Message: commons.String(internal.SuccessMessage),
	}
	return c.JSON(http.StatusOK, response)
}
