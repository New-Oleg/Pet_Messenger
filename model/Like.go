package model

import "time"

type Like struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	UserID    string `gorm:"index;not null"`
	PostID    string `gorm:"index;not null"`
	CreatedAt time.Time
}
