package dto

type ConversationDTO struct {
	ID       string             `json:"id"`
	Messages []DirectMessageDTO `json:"messages,omitempty"`
	Users    []UserResponse     `json:"users,omitempty"`
}
