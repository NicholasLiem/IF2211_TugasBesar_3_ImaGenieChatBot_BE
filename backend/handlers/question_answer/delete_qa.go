package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteQuestionAnswer(c *fiber.Ctx) error {
	var question models.QuestionAnswer
	if err := c.BodyParser(&question); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if question.Question == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Question field is required")
	}

	err := database.DB.Db.Where("question = ?", question.Question).Delete(&models.QuestionAnswer{}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete question answer")
	}

	return c.SendString("Question and answer deleted successfully")
}
