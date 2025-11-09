package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/service"
)

type ConversationController struct {
	service *service.ConversationService
}

func NewConversationController(s *service.ConversationService) *ConversationController {
	return &ConversationController{service: s}
}

func (c *ConversationController) StartConversation(ctx *gin.Context) {
	var req struct {
		TargetUserID string `json:"target_user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetString("userID")
	conv, err := c.service.StartConversation(ctx, userID, req.TargetUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start conversation"})
		return
	}

	ctx.JSON(http.StatusOK, conv)
}

func (c *ConversationController) SendMessage(ctx *gin.Context) {
	var req dto.DirectMessageCreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetString("userID")
	convID := ctx.Param("id")

	msg := &model.DirectMessage{
		ConversationID: convID,
		SenderID:       userID,
		Text:           req.Text,
	}

	if err := c.service.SendMessage(ctx, msg); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	ctx.JSON(http.StatusOK, msg)
}

func (c *ConversationController) GetConversations(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	convs, err := c.service.GetConversations(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get conversations"})
		return
	}
	ctx.JSON(http.StatusOK, convs)
}

func (c *ConversationController) GetMessages(ctx *gin.Context) {
	convID := ctx.Param("id")
	msgs, err := c.service.GetMessages(ctx, convID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}
	ctx.JSON(http.StatusOK, msgs)
}
