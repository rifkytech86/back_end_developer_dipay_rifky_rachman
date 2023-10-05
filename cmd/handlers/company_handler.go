package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) CompanyHandler() controller.ICompanyController {
	table := model.NewCompanies()
	companyRepository := repositories.NewCompanyRepository(h.Application.MongoDBClient, table)
	companyUseCase := usecase.NewCompanyUseCase(companyRepository, h.Application.JWT)
	companyController := controller.NewCompanyController(companyUseCase, h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return companyController
}
