package dto

import "time"

type UserWithConversationsDTO struct {
	ID            string            `json:"id"`
	Username      string            `json:"username"`
	Email         string            `json:"email"`
	Bio           string            `json:"bio,omitempty"`
	AvatarURL     string            `json:"avatar_url,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Conversations []ConversationDTO `json:"conversations,omitempty"`
}
