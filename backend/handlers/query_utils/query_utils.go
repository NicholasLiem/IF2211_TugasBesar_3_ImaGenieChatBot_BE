package query_utils

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllQuestionAnswers() ([]models.QuestionAnswer, error) {
	var qas []models.QuestionAnswer
	if err := database.DB.Db.Find(&qas).Error; err != nil {
		return nil, err
	}
	return qas, nil
}

func InsertMessage(message models.Message) error {
	if err := database.DB.Db.Create(&message).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create message")
	}
	return nil
}
