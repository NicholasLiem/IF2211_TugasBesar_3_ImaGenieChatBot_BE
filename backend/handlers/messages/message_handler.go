package messages

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

func MessageHandler(c *fiber.Ctx) error {
	// Check if session ID is valid
	sessionIDString := c.Params("session_id")
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid session ID")
	}

	// Check if chat session exists
	var existingSession models.ChatSession
	err = database.DB.Db.Where("id = ?", sessionID).First(&existingSession).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Chat session not found, register your chat session!")
	}

	// Parse message from request body
	var message models.Message
	if err := c.BodyParser(&message); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate message
	if message.Text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Text field is required")
	}

	// Insert user message
	userMessage := models.Message{
		SessionID: sessionID,
		Sender:    "user",
		Text:      message.Text,
	}
	if err := query_utils.InsertMessage(userMessage); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to insert user's message to database")
	}

	// Create response message struct
	responseMessage := models.Message{
		SessionID: sessionID,
		Sender:    "bot",
	}

	// Handle QA queries
	message.Text = strings.ToLower(message.Text)
	if isQAQuery(message.Text) {
		result, err := user_query.QuestionAnswerClassifier(message.Text)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		switch result {
		case user_query.SuccessAdd:
			responseMessage.Text = "Question added successfully"
		case user_query.SuccessUpdate:
			responseMessage.Text = "Question updated successfully"
		case user_query.SuccessDelete:
			responseMessage.Text = "Question deleted successfully"
		default:
			return fiber.NewError(fiber.StatusBadRequest, "Invalid query")
		}
	} else {
		// Handle regular queries
		response, err := user_query.QAStringMatchingHandler(message.Text)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		responseMessage.Text = response
	}

	// Insert bot message
	if err := query_utils.InsertMessage(responseMessage); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to insert bot's message to database")
	}

	// Return response message
	return c.Status(fiber.StatusOK).JSON(responseMessage)
}

func isQAQuery(text string) bool {
	r := regexp.MustCompile(`^(tambahkan|add|ubah|update|hapus|delete) pertanyaan (?:(?P<question>.+?)(?: dengan jawaban (?P<answer>.+))?)?$`)
	return r.MatchString(text)
}
