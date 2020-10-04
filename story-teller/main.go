package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "I will /tell you a story..\n")
	})
	e.GET("/tell", TellStory)
	e.Logger.Fatal(e.Start(":4444"))
}

func TellStory(c echo.Context) error {
	story := "It was <TODO add weather here> day.\n" // TODO - add weather from the /weather of world microservice
	story += "And since all good stories have a hero...\n"
	return c.String(http.StatusOK, story)
}
