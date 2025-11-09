package dto

type PostCreateDTO struct {
	Text string `json:"text" binding:"required,min=1,max=1000"`
}
