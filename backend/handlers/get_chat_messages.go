package handlers

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetChatMessages(c *fiber.Ctx) error {
	sessionIDString := c.Params("session_id")
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid session ID")
	}

	var retrievedMessages []models.Message
	database.DB.Db.Order("created_at ASC").Find(&retrievedMessages, "session_id = ?", sessionID.String())
	if len(retrievedMessages) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "No messages found for chat session")
	}

	return c.Status(fiber.StatusOK).JSON(retrievedMessages)
}
