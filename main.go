package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shamith16/mindraeAPI/services/https"
)

func main() {
	app := fiber.New()
	app.Static("/showhome", "scratches/movie-home.json")
	https.MindraeMovieHome()
	app.Listen(":6969")
}
