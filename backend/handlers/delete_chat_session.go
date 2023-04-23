package handlers

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteChatSession(c *fiber.Ctx) error {
	sessionIDString := c.Params("session_id")
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid session ID")
	}

	err = database.DB.Db.Where("id = ?", sessionID).Delete(&models.ChatSession{}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete chat session")
	}

	err = database.DB.Db.Where("session_id = ?", sessionID).Delete(&models.Message{}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete messages for chat session")
	}

	return c.SendString("Chat session and associated messages deleted successfully")
}
