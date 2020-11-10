package main

import (
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" port=5432 user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable TimeZone=Europe/Sofia"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// Create DB schema
	err = db.AutoMigrate(&Story{})
	if err != nil {
		log.Fatalln(err)
	}
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
