package question_answer

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func CreateQuestionAnswer(c *fiber.Ctx) error {
	var newQuestionAnswer models.QuestionAnswer
	if err := c.BodyParser(&newQuestionAnswer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if newQuestionAnswer.Question == "" || newQuestionAnswer.Answer == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Question and answer fields are required")
	}

	if err := database.DB.Db.Create(&newQuestionAnswer).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create question answer")
	}

	response := map[string]interface{}{
		"qid":      newQuestionAnswer.ID,
		"question": newQuestionAnswer.Question,
		"answer":   newQuestionAnswer.Answer,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
