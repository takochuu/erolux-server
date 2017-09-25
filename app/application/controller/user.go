package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type UserController struct{}

func (r UserController) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (r UserController) List(c echo.Context) error {
	return c.String(http.StatusOK, "Fugafuga")
}
