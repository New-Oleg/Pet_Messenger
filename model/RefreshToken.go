package model

import (
	"time"
)

type RefreshToken struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	UserID    string    `gorm:"index;not null"`
	Token     string    `gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
