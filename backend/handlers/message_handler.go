package handlers

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

// GetMessages mengembalikan semua data message dari database
func GetMessages(c *fiber.Ctx) error {
	messages := []models.Message{}
	database.DB.Db.Preload("Session").Find(&messages)
	return c.Status(200).JSON(messages)
}

// CreateMessage membuat data message baru di database
func CreateMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&message)
	return c.Status(200).JSON(message)
}
