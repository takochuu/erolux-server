package routes

import (
	"github.com/labstack/echo"

	"github.com/takochuu/erolux-server/app/application/controller"
)

type Routes struct{}

func NewRoutes() Routes {
	return Routes{}
}

func (r Routes) Add(e *echo.Echo) {
	e.GET("/", controller.IndexController{}.Index)
	e.GET("/hoge", controller.IndexController{}.Hoge)
}
