package handlers

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func GetSessions(c *fiber.Ctx) error {
	sessions := []models.Session{}
	database.DB.Db.Preload("Messages").Find(&sessions)
	return c.Status(200).JSON(sessions)
}

func CreateSession(c *fiber.Ctx) error {
	session := new(models.Session)
	if err := c.BodyParser(session); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&session)
	return c.Status(200).JSON(session)
}
