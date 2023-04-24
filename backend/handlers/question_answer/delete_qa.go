package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/user_query"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteQuestionAnswer(c *fiber.Ctx) error {
	var message models.Message
	if err := c.BodyParser(&message); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if message.Text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Query can't be empty")
	}

	if err := user_query.DeleteQuestionAnswer(message.Text); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendString("Question answer berhasil dihapus")
}
