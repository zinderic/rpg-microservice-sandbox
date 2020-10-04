package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Explore around.. there are /places to go.\n")
	})
	e.GET("/places", func(c echo.Context) error {
		return c.String(http.StatusOK, "TODO - add some interesting places\n")
	})
	e.Logger.Fatal(e.Start(":3333"))
}
