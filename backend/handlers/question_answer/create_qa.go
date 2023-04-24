package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func CreateQuestionAnswer(c *fiber.Ctx) error {
	var message models.Message
	if err := c.BodyParser(&message); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if message.Text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Text field is required")
	}

	if err := user_query.AddOrUpdateQuestionAnswer(message.Text); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendString("Question answer berhasil ditambahkan")
}
