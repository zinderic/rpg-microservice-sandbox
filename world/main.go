package main

import (
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Explore around.. there are /places to go.\n")
	})
	e.GET("/places", func(c echo.Context) error {
		return c.String(http.StatusOK, "TODO - add some interesting places\n")
	})
	e.GET("/weather", Weather)
	e.Logger.Fatal(e.Start(":3333"))
}

func Weather(c echo.Context) error {
	weatherOptions := []string{
		"snowing heavily",
		"raining lightly",
		"sunny",
	}
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	weatherReport := fmt.Sprint(weatherOptions[rand.Intn(len(weatherOptions))], "\n")
	return c.String(http.StatusOK, weatherReport)
}
