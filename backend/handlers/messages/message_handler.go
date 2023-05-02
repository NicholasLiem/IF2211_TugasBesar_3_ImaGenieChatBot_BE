package messages

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/calculator"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/date"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/extra"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	} else if isMathQuery(message.Text) {
		c := &calculator.Calculator{}
		mathRegex := regexp.MustCompile(`^(hitunglah|berapakah)\s(.+)$`)
		match := mathRegex.FindStringSubmatch(message.Text)
		mathExpr := match[2]
		c.InsertInput(mathExpr)
		c.Calculate()
		if c.IsValid() {
			responseMessage.Text = strconv.FormatFloat(c.GetSolution(), 'f', 2, 64)
		} else {
			responseMessage.Text = c.GetErrorMessage()
		}
	} else if isDateQuery(message.Text) {
		d := &date.Date{}
		dateRegex := regexp.MustCompile(`^hari apakah tanggal (\d{1,2}\/\d{1,2}\/\d{4})\?$`)
		match := dateRegex.FindStringSubmatch(message.Text)
		dateString := match[1]
		d.GetDayFromDate(dateString)
		if d.Valid {
			responseMessage.Text = d.GetDateResult()
		} else {
			responseMessage.Text = d.GetErrorMessage()
		}
	} else if isGameQuery(message.Text) {
		rps := &extra.RPSGame{}
		rpsRegex := regexp.MustCompile(`^mainkan suit dengan (\w+)$`)
		match := rpsRegex.FindStringSubmatch(message.Text)
		inputString := match[1]
		rps.PlayGame(inputString)
		responseMessage.Text = rps.GetMessage()
	} else if isRandomPickQuery(message.Text){
		rd := &extra.RandomPick{}
		rdRegex := regexp.MustCompile(`^pilih (\d+) dari ([\w\s]+)$`)
		match := rdRegex.FindStringSubmatch(message.Text)
		amountString := match[1]
		inputString := match[2]
		rd.Pick(amountString, inputString)
		responseMessage.Text = rd.GetMessage()
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
	responseMessage.CreatedAt = time.Now()
	// Return response message
	return c.Status(fiber.StatusOK).JSON(responseMessage)
}

func isQAQuery(text string) bool {
	r := regexp.MustCompile(`^(tambahkan|add|ubah|update|hapus|delete) pertanyaan (?:(?P<question>.+?)(?: dengan jawaban (?P<answer>.+))?)?$`)
	return r.MatchString(text)
}

func isMathQuery(text string) bool {
	r := regexp.MustCompile(`^(hitunglah|berapakah)\s.+[0-9+\-*/().\s]+$`)
	return r.MatchString(text)
}

func isDateQuery(text string) bool {
	r := regexp.MustCompile(`^hari apakah tanggal (\d{1,2}\/\d{1,2}\/\d{4})\?$`)
	return r.MatchString(text)
}

func isGameQuery(text string) bool {
	r := regexp.MustCompile(`^mainkan suit dengan (\w+)$`)
	return r.MatchString(text)
}

func isRandomPickQuery(text string) bool {
	r := regexp.MustCompile(`^pilih (\d+) dari ([\w\s]+)$`)
	return r.MatchString(text)
}

