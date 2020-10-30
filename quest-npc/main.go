package main

import (
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Story struct {
	gorm.Model
	Title string
	Body  string
}

var (
	stories = []Story{
		{Title: "Hello World", Body: "This is a story about a developer who didn't know what to put here."},
		{Title: "Another world", Body: "Yeah.. right."},
	}
)

func main() {

	dsn := "host=quest-npc-db-svc port=5432 user=gorm password=gorm dbname=gorm sslmode=disable TimeZone=Europe/Sofia"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// Create DB schema
	db.AutoMigrate(&Story{})

	// Add the dummy records from the var section
	db.Create(stories)

	// TODO get quest stories from the database instead of those hardcoded values
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Nothing to see here.. move along!\n")
	})
	e.GET("/quests", func(c echo.Context) error {
		return c.String(http.StatusOK, "There will be quests here..\n")
	})
	e.Logger.Fatal(e.Start(":2222"))
}
