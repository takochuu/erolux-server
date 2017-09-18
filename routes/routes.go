package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

type Routes struct{}

func NewRoutes() Routes {
	return Routes{}
}

func (r Routes) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
