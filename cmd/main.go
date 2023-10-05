package main

import (
	"github.com/dipay/api"
	"github.com/dipay/bootstrap"
	"github.com/dipay/cmd/handlers"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	customMiddleware "github.com/dipay/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
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

func authMiddleware(jwt jwt.IJWTRSAToken) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return commons.ErrorResponse(c, http.StatusForbidden, internal.ErrMissingAuthorizationHeader.GetCode(), internal.ErrMissingAuthorizationHeader.String())
			}

			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				return commons.ErrorResponse(c, http.StatusForbidden, internal.ErrInvalidTokenFormat.GetCode(), internal.ErrInvalidTokenFormat.String())
			}

			tokenString := splitToken[1]
			userID, userName, err := jwt.ParserToken(tokenString)
			if err != nil {
				return commons.ErrorResponse(c, http.StatusForbidden, internal.ErrInvalidToken.GetCode(), internal.ErrInvalidToken.String())
			}
			if userID == "" {
				return commons.ErrorResponse(c, http.StatusForbidden, internal.ErrInvalidToken.GetCode(), internal.ErrInvalidToken.String())
			}
			if userName == "" {
				return commons.ErrorResponse(c, http.StatusForbidden, internal.ErrInvalidToken.GetCode(), internal.ErrInvalidToken.String())
			}

			c.Set("userID", userID)
			c.Set("userName", userName)
			return next(c)
		}
	}

}

func setMiddleware(e *echo.Echo, jwt jwt.IJWTRSAToken) *echo.Echo {
	mwRoot := customMiddleware.NewMiddlewareRoot()
	routeGroup := mwRoot.Group("/api")
	routeGroup.Use(authMiddleware(jwt))
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
