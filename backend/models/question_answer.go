package models

type QuestionAnswer struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;" json:"qid"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
