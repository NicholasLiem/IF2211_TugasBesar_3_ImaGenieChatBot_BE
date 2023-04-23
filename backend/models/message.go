package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageID string  `json:"message_id" gorm:"type:text; not null;default:null"`
	Session   Session `json:"session"`
	UserID    uint    `json:"user_id"`
	Content   string  `json:"content"`
	TimeFrame string  `json:"time_frame"`
}
