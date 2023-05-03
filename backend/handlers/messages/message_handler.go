package messages

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/calculator"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/date"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/extra"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
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
		SessionID:   sessionID,
		Sender:      "user",
		PatternType: message.PatternType,
		Text:        message.Text,
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
	userQueries := strings.Split(message.Text, ".")
	if len(userQueries) > 1 {
		// Kalau misalnya dalam kalimat terdiri dari beberapa pertanyaan
		var resultText string
		count := 0
		for index := range userQueries {
			if userQueries[index] != "" {
				resultingText, err := ResponseText(userQueries[index], message.PatternType)
				if err != nil {
					return fiber.NewError(fiber.StatusBadRequest, "Fail to get answer response")
				}
				resultText = resultText + "Answer for Question No." + strconv.Itoa(count+1) + ": \n " + resultingText + "\n\n"
				count++
			}
		}
		// Menghilangkan \n\n dari akhir kalimat
		resultText = strings.TrimSuffix(resultText, "\n\n")
		responseMessage.Text = resultText
	} else {
		// Kalo misalnya kalimat terdiri dari 1 kalimat saja.
		responseMessage.Text, err = ResponseText(message.Text, message.PatternType)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Fail to get answer response")
		}
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
	r := regexp.MustCompile(`^[\s]*(tambahkan|add|ubah|update|hapus|delete)[\s]+pertanyaan[\s]+(?:(?P<question>.+?)(?:[\s]+dengan[\s]+jawaban[\s]+(?P<answer>.+))?)?[\s]*$`)
	return r.MatchString(text)
}

func isMathQuery(text string) bool {
	r := regexp.MustCompile(`^[\s]*(hitunglah|berapakah)[\s]+.+[0-9+\-*/().\s]+[\s]*$`)
	return r.MatchString(text)
}

//func isSpaceQuery(text string) bool {
//	r := regexp.MustCompile(`(\s)+`)
//	return r.MatchString(text)
//}

func isDateQuery(text string) bool {
	r := regexp.MustCompile(`^[\s]*hari[\s]+apakah[\s]+tanggal[\s]+(\d{1,2}\/\d{1,2}\/[\d]+)[\?]*[\s]*$`)
	return r.MatchString(text)
}

func isGameQuery(text string) bool {
	r := regexp.MustCompile(`^[\s]*mainkan[\s]+suit[\s]+dengan[\s]+(\w+)[\s]*$`)
	return r.MatchString(text)
}

func isRandomPickQuery(text string) bool {
	r := regexp.MustCompile(`^[\s]*pilih[\s]+(\d+)[\s]+dari[\s]+([\w\s]+)[\s]*$`)
	return r.MatchString(text)
}

func ResponseText(text string, patternType string) (string, error) {
	var response string
	if isQAQuery(text) {
		result, err := user_query.QuestionAnswerClassifier(text)
		if err != nil {
			return response, fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		switch result {
		case user_query.SuccessAdd:
			response = "Question added successfully"
		case user_query.SuccessUpdate:
			response = "Question updated successfully"
		case user_query.SuccessDelete:
			response = "Question deleted successfully"
		default:
			return "", fiber.NewError(fiber.StatusBadRequest, "Invalid query")
		}
	} else if isMathQuery(text) {
		c := &calculator.Calculator{}
		mathRegex := regexp.MustCompile(`^[\s]*(hitunglah|berapakah)\s(.+)[\s]*$`)
		match := mathRegex.FindStringSubmatch(text)
		mathExpr := match[2]
		c.InsertInput(mathExpr)
		c.Calculate()
		if c.IsValid() {
			response = strconv.FormatFloat(c.GetSolution(), 'f', 2, 64)
		} else {
			response = c.GetErrorMessage()
		}
	} else if isDateQuery(text) {
		d := &date.Date{}
		dateRegex := regexp.MustCompile(`^[\s]*hari[\s]+apakah[\s]+tanggal[\s]+(\d{1,2}\/\d{1,2}\/[\d]+)[\?]*[\s]*$`)
		match := dateRegex.FindStringSubmatch(text)
		dateString := match[1]
		d.GetDayFromDate(dateString)
		if d.Valid {
			response = d.GetDateResult()
		} else {
			response = d.GetErrorMessage()
		}
	} else if isGameQuery(text) {
		rps := &extra.RPSGame{}
		rpsRegex := regexp.MustCompile(`^[\s]*mainkan[\s]+suit[\s]+dengan[\s]+(\w+)[\s]*$`)
		match := rpsRegex.FindStringSubmatch(text)
		inputString := match[1]
		rps.PlayGame(inputString)
		response = rps.GetMessage()
	} else if isRandomPickQuery(text) {
		rd := &extra.RandomPick{}
		rdRegex := regexp.MustCompile(`^[\s]*pilih[\s]+(\d+)[\s]+dari[\s]+([\w\s]+)[\s]*$`)
		match := rdRegex.FindStringSubmatch(text)
		amountString := match[1]
		inputString := match[2]
		rd.Pick(amountString, inputString)
		response = rd.GetMessage()
	} else {
		// Handle regular queries
		queryResponse, err := user_query.QAStringMatchingHandler(text, patternType)
		if err != nil {
			return "", fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		response = queryResponse
	}
	return response, nil
}
