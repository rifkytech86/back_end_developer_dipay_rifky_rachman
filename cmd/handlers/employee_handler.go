package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/model"
	"github.com/dipay/pkg/httpClient"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) EmployeeHandler() controller.IEmployeeController {
	table := model.NewEmployees()
	tableCompany := model.NewCompanies()
	initialHttpClient := httpClient.NewClient()
	employeeRepository := repositories.NewEmployeeRepository(h.Application.MongoDBClient, table, initialHttpClient, h.Application.ENV.EmailService)
	companyRepository := repositories.NewCompanyRepository(h.Application.MongoDBClient, tableCompany)

	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository, companyRepository)
	employeeController := controller.NewEmployeeController(employeeUseCase, h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return employeeController
}
