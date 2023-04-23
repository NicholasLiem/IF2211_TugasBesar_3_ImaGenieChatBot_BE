package models

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SessionId string `json:"session_id" gorm:"type:text; not null;default:null"`
	Title     string `json:"title" gorm:"text;not null;default:null"`
}
