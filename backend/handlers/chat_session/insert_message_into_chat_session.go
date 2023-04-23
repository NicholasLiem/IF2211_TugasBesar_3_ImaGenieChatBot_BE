package chat_session

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func InsertMessageToChatSession(c *fiber.Ctx) error {
	var newMessage models.Message
	if err := c.BodyParser(&newMessage); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	sessionIDString := c.Params("session_id")
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid session ID")
	}

	var existingSession models.ChatSession
	err = database.DB.Db.Where("id = ?", sessionID).First(&existingSession).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Chat session not found")
	}

	newMessage.SessionID = sessionID
	if err := database.DB.Db.Create(&newMessage).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create message")
	}

	return c.Status(fiber.StatusCreated).JSON(newMessage)
}
