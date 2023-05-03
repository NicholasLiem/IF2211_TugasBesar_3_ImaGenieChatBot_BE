package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	SessionID   uuid.UUID `gorm:"type:uuid;index;" json:"sessionId"`
	Sender      string    `json:"sender"`
	Text        string    `json:"text"`
	PatternType string    `json:"patternType"`
	CreatedAt   time.Time `json:"createdAt"`
}
