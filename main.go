package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/shamith16/mindraeAPI/services/https"
)

func main() {
	app := fiber.New()
	https.MindraeMovieHome()
	c := cron.New()
	c.AddFunc("90 * * * *", func() {
		fmt.Println("Cron job started")
		https.MindraeMovieHome()
	})

	c.Start()
	app.Static("/showhome", "movie-home.json")
	app.Listen(":6969")

}
