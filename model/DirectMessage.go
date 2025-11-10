package model

import "time"

type DirectMessage struct {
	ID             string    `gorm:"type:char(36);primaryKey" json:"id"`
	ConversationID string    `gorm:"not null;index" json:"conversation_id"`
	SenderID       string    `gorm:"not null;index" json:"sender_id"`
	Text           string    `gorm:"type:text;not null" json:"text"`
	CreatedAt      time.Time `json:"created_at"`
}
