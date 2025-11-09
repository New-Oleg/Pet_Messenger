package dto

import "time"

type CommentResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
