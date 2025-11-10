package dto

type StartConversationDTO struct {
	TargetUserID string `json:"target_user_id" binding:"required"`
}
