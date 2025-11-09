package model

import (
	"time"
)

// User — сущность пользователя
type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Bio       string    `gorm:"type:text" json:"bio,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// связи (один-ко-многим)
	Posts    []Post    `gorm:"foreignKey:AuthorID" json:"-"`
	Likes    []Like    `gorm:"foreignKey:UserID" json:"-"`
	Messages []Comment `gorm:"foreignKey:SenderID" json:"-"`
}
