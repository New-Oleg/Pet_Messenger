package dto

import "time"

type PostResponse struct {
	ID         string    `json:"id"`
	AuthorID   string    `json:"author_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LikesCount int64     `json:"lokesCount" `
}
