package api

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewAPIHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
