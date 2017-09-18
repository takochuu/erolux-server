package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type IndexController struct{}

func (r IndexController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (r IndexController) Hoge(c echo.Context) error {
	return c.String(http.StatusOK, "Fugafuga")
}
