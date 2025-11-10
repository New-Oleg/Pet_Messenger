package model

import (
	"time"
)

type Post struct {
	ID         string `gorm:"type:char(36);primaryKey" json:"id"`
	AuthorID   string `gorm:"index;not null"`
	Text       string `gorm:"text;not null"`
	LikesCount int64  `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
