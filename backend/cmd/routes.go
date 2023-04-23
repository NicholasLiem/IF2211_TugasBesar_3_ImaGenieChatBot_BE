package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
}
