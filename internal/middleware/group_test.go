package middleware

import (
	"github.com/dipay/internal/jwt/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newMockMiddlewareRoot() *middlewareRoot {
	e := echo.New()
	return &middlewareRoot{
		middlewares: map[string]map[string][]echo.MiddlewareFunc{},
		router:      echo.NewRouter(e),
		echo:        e,
	}
}

func TestGroup(t *testing.T) {
	mwr := newMockMiddlewareRoot()
	mockJWT := new(mocks.IJWTRSAToken)
	g := mwr.Group("/api", AuthMiddleware(mockJWT), AuthMiddleware(mockJWT))
	assert.Equal(t, "/api", g.prefix)
}

func TestCONNECT(t *testing.T) {
	mockJWT := new(mocks.IJWTRSAToken)
	mock1 := AuthMiddleware(mockJWT)
	mock2 := AuthMiddleware(mockJWT) // Create a second instance

	mwr := newMockMiddlewareRoot()
	g := mwr.Group("/api")
	g.CONNECT("/users", mock1, mock2)

}
