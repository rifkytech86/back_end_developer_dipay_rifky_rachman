package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) EmployeeHandler() controller.IEmployeeController {
	table := model.NewEmployees()
	tableCompany := model.NewCompanies()
	employeeRepository := repositories.NewEmployeeRepository(h.Application.MongoDBClient, table)
	companyRepository := repositories.NewCompanyRepository(h.Application.MongoDBClient, tableCompany)

	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository, companyRepository, h.Application.JWT)
	employeeController := controller.NewEmployeeController(employeeUseCase, h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return employeeController
}
