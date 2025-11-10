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

// POST /conversations
// @Summary Start conversation
// @Description Create a new conversation with a target user
// @Tags conversations
// @Accept json
// @Produce json
// // @Param body body dto.StartConversationDTO true "Target user ID"
// @Success 200 {object} model.Conversation
// @Router /conversations [post]
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

// POST /conversations/:id/messages
// @Summary Send message
// @Description Send a message in a conversation
// @Tags conversations
// @Accept json
// @Produce json
// @Param id path string true "Conversation ID"
// @Param message body dto.DirectMessageCreateDTO true "Message content"
// @Success 200 {object} model.DirectMessage
// @Router /conversations/{id}/messages [post]
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

// GET /conversations
// @Summary List conversations
// @Description Get all conversations of authenticated user
// @Tags conversations
// @Produce json
// @Success 200 {array} model.Conversation
// @Router /conversations [get]
func (c *ConversationController) GetConversations(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	convs, err := c.service.GetConversations(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get conversations"})
		return
	}
	ctx.JSON(http.StatusOK, convs)
}

// GET /conversations/:id/messages
// @Summary Get messages
// @Description Get all messages in a conversation
// @Tags conversations
// @Produce json
// @Param id path string true "Conversation ID"
// @Success 200 {array} model.DirectMessage
// @Router /conversations/{id}/messages [get]
func (c *ConversationController) GetMessages(ctx *gin.Context) {
	convID := ctx.Param("id")
	msgs, err := c.service.GetMessages(ctx, convID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}
	ctx.JSON(http.StatusOK, msgs)
}
