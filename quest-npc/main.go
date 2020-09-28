package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Nothing to see here.. move along!\n")
	})
	e.GET("/quests", func(c echo.Context) error {
		return c.String(http.StatusOK, "There will be quests here..\n")
	})
	e.Logger.Fatal(e.Start(":2222"))
}
