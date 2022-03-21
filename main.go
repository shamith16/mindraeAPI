package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/shamith16/mindraeAPI/services/https"
	"os"
)

const PORT = "3000"

func main() {
	app := fiber.New()
	port := getEnv("PORT", PORT)
	https.MindraeMovieHome()
	c := cron.New()
	c.AddFunc("90 * * * *", func() {
		fmt.Println("Cron job started")
		https.MindraeMovieHome()
	})

	c.Start()
	app.Static("/showhome", "movie-home.json")
	app.Listen(fmt.Sprintf(":%s", port))
}

// Gets default value passed if no value exist for given environment variable.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
