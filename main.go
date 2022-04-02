package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/shamith16/mindraeAPI/services/fetch"
	"github.com/shamith16/mindraeAPI/utils"
	"os"
	"time"
)

const PORT = "3000"
const MIN = "360"

//TODO Add Api endpoint to change cron via post

func init() {
	os.Setenv("TZ", "Asia/Kolkata")
	time.AfterFunc(3*time.Minute, func() {
		_ = utils.WriteToFile("logs.txt", []byte("Running MovieHome() after 3 minutes"), "logs", "append")
		fetch.MovieHome()

	})
}

func main() {
	app := fiber.New()
	port := getEnv("PORT", PORT)

	c := cron.New()

	c.AddFunc(MIN+" * * * *", func() {
		fmt.Println("Cron job started")
		fetch.MovieHome()
	})

	c.Start()

	app.Static("/moviehome", "jsons/movie-home.json")
	app.Static("/logs", "logs/logs.txt")

	appPort := fmt.Sprintf(":%s", port)
	_ = utils.WriteToFile("logs.txt", []byte(appPort), "logs", "append")
	_ = app.Listen(appPort)
}

// Gets default value passed if no value exist for given environment variable.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
