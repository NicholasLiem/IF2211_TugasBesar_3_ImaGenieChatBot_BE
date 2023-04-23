package models

import (
	"github.com/google/uuid"
	"time"
)

type ChatSession struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	Messages  []Message `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE;" json:"messages"`
	CreatedAt time.Time `json:"createdAt"`
}
