package handlers

import (
	"github.com/dipay/api"
	"github.com/dipay/bootstrap"
	"github.com/dipay/controller"
	"github.com/labstack/echo/v4"
)

type MyHandler struct {
	Application bootstrap.Application
}

func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

type ServerInterfaceWrapper struct {
	UserAdminHandler        controller.IUserAdminController
	CompanyHandler          controller.ICompanyController
	EmployeeController      controller.IEmployeeController
	CountriesController     controller.ICountriesController
	DuplicateZeroController controller.IDuplicateZeroController
}

func (w *ServerInterfaceWrapper) Hello(ctx echo.Context, param api.HelloParams) error {
	return w.UserAdminHandler.Hello(ctx, param)
}

func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	return w.UserAdminHandler.Login(ctx)
}

func (w *ServerInterfaceWrapper) AddCompany(ctx echo.Context) error {
	return w.CompanyHandler.AddCompany(ctx)
}

func (w *ServerInterfaceWrapper) GetCompany(ctx echo.Context) error {
	return w.CompanyHandler.GetCompany(ctx)
}

func (w *ServerInterfaceWrapper) UpdateCompanyStatusActive(ctx echo.Context, id string) error {
	return w.CompanyHandler.UpdateCompanyStatusActive(ctx, id)
}

func (w *ServerInterfaceWrapper) AddEmployee(ctx echo.Context, id string) error {
	return w.EmployeeController.AddEmployee(ctx, id)
}

func (w *ServerInterfaceWrapper) GetEmployeeByID(ctx echo.Context, id string) error {
	return w.EmployeeController.GetEmployeeByID(ctx, id)
}

func (w *ServerInterfaceWrapper) GetEmployeeByCompanyID(ctx echo.Context, id string) error {
	return w.EmployeeController.GetEmployeeByCompanyID(ctx, id)
}

func (w *ServerInterfaceWrapper) UpdateEmployeeData(ctx echo.Context, companyId string, employeeId string) error {
	return w.EmployeeController.UpdateEmployeeData(ctx, companyId, employeeId)
}

func (w *ServerInterfaceWrapper) DeleteEmployeeByID(ctx echo.Context, id string) error {
	return w.EmployeeController.DeleteEmployeeByID(ctx, id)
}

func (w *ServerInterfaceWrapper) GetDataCountries(ctx echo.Context) error {
	return w.CountriesController.GetDataCountries(ctx)
}

func (w *ServerInterfaceWrapper) DuplicateZero(ctx echo.Context) error {
	return w.DuplicateZeroController.DuplicateZero(ctx)
}
