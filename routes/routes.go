package routes

import (
	"net/http"
	"os"

	"encoding/json"

	"github.com/labstack/echo"
)

type Routes struct{}

func NewRoutes() Routes {
	return Routes{}
}

type Route struct {
	Method string           `json: "method"`
	Path   string           `json: "path"`
	Name   echo.HandlerFunc `json: "name"`
}

func (r Routes) Get() []Route {
	file, err := os.Open("./routes/routes.json")
	if err != nil {
		// TODO Error Handling
	}
	defer file.Close()

	var d []Route
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&d); err != nil {
		// TODO Error Handling
	}
	return d
}

func (r Routes) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
