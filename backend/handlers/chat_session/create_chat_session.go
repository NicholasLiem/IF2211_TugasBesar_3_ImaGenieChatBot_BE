package chat_session

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func CreateChatSession(c *fiber.Ctx) error {
	var newMessage models.Message
	if err := c.BodyParser(&newMessage); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	sessionIDString := c.Params("session_id")
	var sessionID uuid.UUID
	if sessionIDString == "" {
		sessionID = uuid.New()
	} else {
		var err error
		sessionID, err = uuid.Parse(sessionIDString)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid session ID")
		}
	}

	var existingSession models.ChatSession
	err := database.DB.Db.Where("id = ?", sessionID).First(&existingSession).Error
	if err != nil {
		newChatSession := models.ChatSession{
			CreatedAt: time.Now(),
			ID:        sessionID,
		}

		if err := database.DB.Db.Create(&newChatSession).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create chat session")
		}

		existingSession = newChatSession
	}

	newMessage.SessionID = sessionID
	if err := database.DB.Db.Create(&newMessage).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create message")
	}

	response := map[string]interface{}{
		"session_id": existingSession.ID,
		"messages":   []models.Message{newMessage},
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
