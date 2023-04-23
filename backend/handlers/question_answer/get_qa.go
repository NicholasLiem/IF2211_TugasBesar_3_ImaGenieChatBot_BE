package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func GetQuestionAnswers(c *fiber.Ctx) error {
	var questionAnswers []models.QuestionAnswer
	err := database.DB.Db.Find(&questionAnswers).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve question answers")
	}

	if len(questionAnswers) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "No question answers found in the database")
	}

	return c.Status(fiber.StatusOK).JSON(questionAnswers)
}
