package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetMessages)
	app.Get("/getSession", handlers.GetSessions)
	app.Post("/createsession", handlers.CreateSession)
	app.Post("/createmessage", handlers.CreateMessage)
}
