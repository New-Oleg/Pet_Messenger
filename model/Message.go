package model

import "time"

type Message struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string `gorm:"index;not null"`
	PostID    string `gorm:"index;not null"`
	Text      string `gorm:"index;not null"`
	CreatedAt time.Time
}
