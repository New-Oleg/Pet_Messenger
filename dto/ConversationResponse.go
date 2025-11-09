package dto

import "time"

type ConversationResponse struct {
	ID        string                  `json:"id"`
	User1ID   string                  `json:"user1_id"`
	User2ID   string                  `json:"user2_id"`
	Messages  []DirectMessageResponse `json:"messages,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}
