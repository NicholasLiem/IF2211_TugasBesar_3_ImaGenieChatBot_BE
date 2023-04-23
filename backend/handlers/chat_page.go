package handlers

import "github.com/gofiber/fiber/v2"

func ChatPage(c *fiber.Ctx) error {
	return c.SendString("Received POST request to /chat/ route")
}
