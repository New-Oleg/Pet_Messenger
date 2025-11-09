package dto

import "time"

type DirectMessageDTO struct {
	ID        string    `json:"id"`
	SenderID  string    `json:"sender_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
