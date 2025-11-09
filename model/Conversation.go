package model

import "time"

// Conversation — диалог между двумя пользователями
type Conversation struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserAID   string    `gorm:"not null" json:"user_a_id"`
	UserBID   string    `gorm:"not null" json:"user_b_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Messages []DirectMessage `gorm:"foreignKey:ConversationID" json:"messages,omitempty"`
}
