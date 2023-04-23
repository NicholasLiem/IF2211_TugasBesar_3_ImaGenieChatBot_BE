package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/chat/", handlers.ChatPage)
	app.Get("/h/:session_id", handlers.GetChatMessages)
	app.Post("/c/:session_id", handlers.CreateChatSession)
}
