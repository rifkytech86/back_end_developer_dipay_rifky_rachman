package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/pkg/httpClient"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) CountriesHandler() controller.ICountriesController {
	initialHttpClient := httpClient.NewClient()
	countriesRepository := repositories.NewCountriesRepository(initialHttpClient)
	countriesUseCase := usecase.NewCountriesUseCase(countriesRepository)
	countriesController := controller.NewCountriesController(countriesUseCase, h.Application.ENV.ContextTimeOut)
	return countriesController
}
