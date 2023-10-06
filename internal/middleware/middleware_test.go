package middleware

import (
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_middlewareRoot_Group(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func Test_middlewareRoot_Group2(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "myAccessToken")

	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func Test_middlewareRoot_Group3(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mockJWT.On("ParserToken", mock.Anything).Return("1", "name", errors.New(internal.ErrorInvalidRequest.String()))
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer myAccessToken")

	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func Test_middlewareRoot_Group31(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mockJWT.On("ParserToken", mock.Anything).Return("", "name", nil)
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer myAccessToken")

	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}
func Test_middlewareRoot_Group32(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mockJWT.On("ParserToken", mock.Anything).Return("1231", "", nil)
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer myAccessToken")

	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func Test_middlewareRoot_Group4(t *testing.T) {
	mwr := NewMiddlewareRoot()

	mockJWT := new(mocks.IJWTRSAToken)
	mockJWT.On("ParserToken", mock.Anything).Return("1", "name", nil)
	mwr.add(http.MethodGet, "/users", AuthMiddleware(mockJWT))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer myAccessToken")

	rec := httptest.NewRecorder()
	c := mwr.echo.NewContext(req, rec)
	h := mwr.Exec(func(c echo.Context) error { return c.String(http.StatusForbidden, "OK") })
	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}
