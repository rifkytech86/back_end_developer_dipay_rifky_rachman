package main

import (
	"github.com/dipay/api"
	"github.com/dipay/bootstrap"
	"github.com/dipay/cmd/handlers"
	"github.com/dipay/internal/jwt"
	customMiddleware "github.com/dipay/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Group("/api")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	app := bootstrap.NewInitializeBootstrap()

	// initial service dependencies
	serve := handlers.NewServiceInitial(app)
	userAdminController := serve.UserAdminHandler()
	companyController := serve.CompanyHandler()
	employeeController := serve.EmployeeHandler()
	countriesController := serve.CountriesHandler()
	duplicateZeroController := serve.DuplicateZeroHandler()

	wrapper := &handlers.ServerInterfaceWrapper{
		UserAdminHandler:        userAdminController,
		CompanyHandler:          companyController,
		EmployeeController:      employeeController,
		CountriesController:     countriesController,
		DuplicateZeroController: duplicateZeroController,
	}
	e = setMiddleware(e, app.JWT)

	api.RegisterHandlers(e, wrapper)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func setMiddleware(e *echo.Echo, jwt jwt.IJWTRSAToken) *echo.Echo {
	mwRoot := customMiddleware.NewMiddlewareRoot()
	routeGroup := mwRoot.Group("/api")
	routeGroup.Use(customMiddleware.AuthMiddleware(jwt))
	{
		routeGroup.POST("/companies")
		routeGroup.GET("/companies")
		routeGroup.PUT("/companies/:id/set_active")
		routeGroup.GET("/employees/:id")
		routeGroup.GET("/companies/:id/employees")
		routeGroup.PUT("/companies/:id/employees/:id")
		routeGroup.GET("/countries")
	}
	e.Use(mwRoot.Exec)
	return e
}
