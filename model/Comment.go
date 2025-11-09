package model

import "time"

type Comment struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string `gorm:"index;not null"`
	PostID    string `gorm:"index;not null"`
	Text      string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
