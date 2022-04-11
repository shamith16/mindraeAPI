package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/shamith16/mindraeAPI/services/fetch/mindraeapi"
	"github.com/shamith16/mindraeAPI/utils"
	"os"
	"time"
)

const PORT = "3000"
const MIN = "360"

//TODO Add Api endpoint to change cron via post

func init() {
	_ = os.Setenv("TZ", "Asia/Kolkata")
	time.AfterFunc(3*time.Minute, func() {
		utils.Logger("Running Home() after 3 minutes\n", "main.go")
		mindraeapi.MovieHome()
	})
}

func main() {
	app := fiber.New()
	port := getEnv("PORT", PORT)

	c := cron.New()

	_ = c.AddFunc(MIN+" * * * *", func() {
		fmt.Println("Cron job started")
		mindraeapi.MovieHome()
	})

	c.Start()

	app.Static("/moviehome", "json/movie-home.json")

	app.Static("/log", "./log", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Download:  true,
		MaxAge:    3600,
		Next:      nil,
	})

	appPort := fmt.Sprintf(":%s", port)

	utils.Logger("Server is running at port"+appPort+"\n", "main.go")

	_ = app.Listen(appPort)
}

// Gets default value passed if no value exist for given environment variable.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
