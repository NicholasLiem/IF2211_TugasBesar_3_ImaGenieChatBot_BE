package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	

	setupRoutes(app)

	app.Listen(":5000")
}
