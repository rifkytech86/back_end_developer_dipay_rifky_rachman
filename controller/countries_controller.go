package controller

import (
	"context"
	"github.com/dipay/api"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type countriesController struct {
	CountriesUseCase usecase.ICountriesUseCase
	ContextTimeOut   int
}

type ICountriesController interface {
	GetDataCountries(ctx echo.Context) error
}

func NewCountriesController(countriesUseCase usecase.ICountriesUseCase, contextTimeOut int) ICountriesController {
	return &countriesController{
		ContextTimeOut:   contextTimeOut,
		CountriesUseCase: countriesUseCase,
	}
}

func (co *countriesController) GetDataCountries(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(co.ContextTimeOut)*time.Second)
	defer cancel()

	listCountries, err := co.CountriesUseCase.GetCountries(ctx)
	if err != nil {
		return err
	}

	var convertedRows []struct {
		Name      *string   `json:"name,omitempty"`
		Region    *string   `json:"region,omitempty"`
		Timezones *[]string `json:"timezones,omitempty"`
	}
	for _, country := range listCountries {
		convertedRows = append(convertedRows, struct {
			Name      *string   `json:"name,omitempty"`
			Region    *string   `json:"region,omitempty"`
			Timezones *[]string `json:"timezones,omitempty"`
		}{
			Name:      &country.Name,
			Region:    &country.Region,
			Timezones: &country.Timezones,
		})
	}

	response := api.GetDataCountriesResponse{
		Code:    commons.String(strconv.Itoa(http.StatusOK)),
		Status:  commons.Int(http.StatusOK),
		Data:    &convertedRows,
		Message: commons.String(internal.SuccessMessage),
	}

	return c.JSON(http.StatusOK, response)
}
