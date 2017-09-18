package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/takochuu/erolux-server/routes"
)

func init() {}

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())

	// routing
	r := routes.NewRoutes()
	r.Add(e)

	e.Logger.Fatal(e.Start(":1323"))
}
