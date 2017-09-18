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
	routes := r.Get()
	for _, v := range routes {
		switch v.Method {
		case "POST":
			e.POST(v.Path, v.Name)
		case "GET":
			e.GET(v.Path, v.Name)
		case "PUT":
			e.PUT(v.Path, v.Name)
		case "DELETE":
			e.DELETE(v.Path, v.Name)
		}

	}
	e.GET("/", r.Index)

	e.Logger.Fatal(e.Start(":1323"))
}
