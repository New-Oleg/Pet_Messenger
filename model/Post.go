package model

import (
	"time"
)

type Post struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	AuthorID   string `gorm:"index;not null"`
	Text       string `gorm:"text;not null"`
	LikesCount int64  `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
