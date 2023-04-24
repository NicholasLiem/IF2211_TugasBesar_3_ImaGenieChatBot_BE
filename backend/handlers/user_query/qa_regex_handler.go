package user_query

import (
	"errors"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/database"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

func AddOrUpdateQuestionAnswer(query string) error {
	r := regexp.MustCompile(`^(Tambahkan|Add|Ubah|Update) pertanyaan (?:(?P<question>.+?)(?: dengan jawaban (?P<answer>.+))?)?$`)
	match := r.FindStringSubmatch(strings.TrimSpace(query))

	if len(match) < 1 {
		return errors.New("invalid query")
	}

	question := strings.TrimSpace(match[2])
	answer := strings.TrimSpace(match[3])

	existingQuestionAnswer := models.QuestionAnswer{}
	if question != "" {
		if err := database.DB.Db.Where("question = ?", question).First(&existingQuestionAnswer).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	switch strings.TrimSpace(match[1]) {
	case "Tambahkan", "Add":
		if question == "" || answer == "" {
			return errors.New("question and answer fields are required")
		}
		if existingQuestionAnswer.ID != 0 {
			existingQuestionAnswer.Answer = answer
			if err := database.DB.Db.Save(&existingQuestionAnswer).Error; err != nil {
				return err
			}
		} else {
			newQuestionAnswer := models.QuestionAnswer{
				Question: question,
				Answer:   answer,
			}
			if err := database.DB.Db.Create(&newQuestionAnswer).Error; err != nil {
				return err
			}
		}
	case "Ubah", "Update":
		if existingQuestionAnswer.ID == 0 {
			return errors.New("question not found")
		}
		existingQuestionAnswer.Answer = answer
		if err := database.DB.Db.Save(&existingQuestionAnswer).Error; err != nil {
			return err
		}
	default:
		return errors.New("invalid query")
	}

	return nil
}

func DeleteQuestionAnswer(query string) error {
	r := regexp.MustCompile(`^(Hapus|Remove) pertanyaan (?P<question>.+)$`)
	match := r.FindStringSubmatch(strings.TrimSpace(query))

	if len(match) > 0 {
		if strings.TrimSpace(match[1]) == "Hapus" || strings.TrimSpace(match[1]) == "Remove" {
			question := strings.TrimSpace(match[2])
			if question == "" {
				return errors.New("question field is required")
			}
			if err := database.DB.Db.Where("question = ?", question).Delete(&models.QuestionAnswer{}).Error; err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("invalid query")
}
