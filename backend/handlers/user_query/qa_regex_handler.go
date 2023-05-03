package user_query

import (
	"errors"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

const (
	SuccessAdd    = 1
	SuccessUpdate = 2
	SuccessDelete = 3
)

func QuestionAnswerClassifier(query string) (int, error) {
	r := regexp.MustCompile(`^(tambahkan|add|ubah|update|hapus|delete) pertanyaan (?:(?P<question>.+?)(?: dengan jawaban (?P<answer>.+))?)?$`)
	match := r.FindStringSubmatch(strings.TrimSpace(query))

	if len(match) > 0 {
		question := strings.TrimSpace(match[2])
		answer := strings.TrimSpace(match[3])

		existingQuestionAnswer := models.QuestionAnswer{}
		if question != "" && answer != "" {
			if err := database.DB.Db.Where("question = ?", question).First(&existingQuestionAnswer).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, err
			}
		} else {
			return 0, errors.New("invalid query")
		}

		if strings.TrimSpace(match[1]) == "tambahkan" || strings.TrimSpace(match[1]) == "add" {
			newQuestionAnswer := models.QuestionAnswer{
				Question: question,
				Answer:   answer,
			}
			if newQuestionAnswer.Question == "" || newQuestionAnswer.Answer == "" {
				return 0, errors.New("question and answer fields are required")
			}
			if existingQuestionAnswer.ID != 0 {
				existingQuestionAnswer.Answer = newQuestionAnswer.Answer
				if err := database.DB.Db.Save(&existingQuestionAnswer).Error; err != nil {
					return 0, err
				}
				return SuccessUpdate, nil
			} else {
				if err := database.DB.Db.Create(&newQuestionAnswer).Error; err != nil {
					return 0, err
				}
				return SuccessAdd, nil
			}
		} else if strings.TrimSpace(match[1]) == "ubah" || strings.TrimSpace(match[1]) == "update" {
			if existingQuestionAnswer.ID == 0 {
				return 0, errors.New("question not found")
			}
			existingQuestionAnswer.Answer = answer
			if err := database.DB.Db.Save(&existingQuestionAnswer).Error; err != nil {
				return 0, err
			}
			return SuccessUpdate, nil
		} else if strings.TrimSpace(match[1]) == "hapus" || strings.TrimSpace(match[1]) == "delete" {
			if existingQuestionAnswer.ID == 0 {
				return 0, errors.New("question not found")
			}
			if err := database.DB.Db.Delete(&existingQuestionAnswer).Error; err != nil {
				return 0, err
			}
			return SuccessDelete, nil
		}
	}
	return 0, errors.New("invalid query")
}
