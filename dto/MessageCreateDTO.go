package dto

type MessageCreateDTO struct {
	ReceiverID string `json:"receiver_id" binding:"required,uuid"`
	Text       string `json:"text" binding:"required,min=1,max=1000"`
}
