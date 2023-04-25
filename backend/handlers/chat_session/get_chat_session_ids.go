package chat_session

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetChatSessionIDs(c *fiber.Ctx) error {
	var chatSessions []models.ChatSession
	if err := database.DB.Db.Model(&models.ChatSession{}).Select("id").Find(&chatSessions).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get chat sessions")
	}

	var sessionIDs []uuid.UUID
	for _, chatSession := range chatSessions {
		sessionIDs = append(sessionIDs, chatSession.ID)
	}

	response := map[string]interface{}{
		"session_ids": sessionIDs,
	}

	return c.JSON(response)
}

func GetChatSessions(c *fiber.Ctx) error{
	var chatSessions []models.ChatSession
	err := database.DB.Db.Find(&chatSessions).Error
	if err != nil{
		return fiber.NewError(fiber.StatusInternalServerError,"Failed to retrieve chat sessions")
	}
	if len(chatSessions) == 0 {
		return fiber.NewError(fiber.StatusNotFound,"No chat sessions found in the database")
	}

	return c.Status(fiber.StatusOK).JSON(chatSessions)
}