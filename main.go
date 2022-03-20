package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/shamith16/mindraeAPI/services/https"
)

func main() {
	https.MindraeMovieHome()
	app := fiber.New()
	https.MindraeMovieHome()
	c := cron.New()
	_, _ = c.AddFunc("90 * * * *", func() {
		fmt.Println("Cron job started")
		https.MindraeMovieHome()
	})
	c.Start()
	app.Static("/showhome", "scratches/movie-home.json")
	app.Listen(":6969")

}
