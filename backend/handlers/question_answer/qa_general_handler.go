package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

const (
	SuccessAdd    = 1
	SuccessUpdate = 2
	SuccessDelete = 3
)

func QuestionAnswerHandler(c *fiber.Ctx) error {
	var message models.Message
	if err := c.BodyParser(&message); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if message.Text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Text field is required")
	}

	status, err := user_query.QuestionAnswerClassifier(message.Text)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var response string
	switch status {
	case SuccessAdd:
		response = "QA berhasil ditambahkan"
	case SuccessUpdate:
		response = "QA berhasil diubah"
	case SuccessDelete:
		response = "QA berhasil dihapus"
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid status code")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": response,
	})
}
