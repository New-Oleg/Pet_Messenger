package dto

import "time"

type DirectMessageResponse struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
}
