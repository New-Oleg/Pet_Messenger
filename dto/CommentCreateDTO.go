package dto

type CommentCreateDTO struct {
	Text string `json:"text" binding:"required,min=1,max=1000"`
}
