package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/service"
)

type MessageController struct {
	messageService *service.MessageService
}

// Конструктор
func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

// POST /posts/:id/comments
func (c *MessageController) CreateComment(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetString("userID")

	var req dto.MessageCreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := c.messageService.CreateComment(ctx, userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, msg)
}

// GET /posts/:id/comments
func (c *MessageController) GetCommentsByPost(ctx *gin.Context) {
	postID := ctx.Param("id")

	comments, err := c.messageService.GetCommentsByPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// DELETE /comments/:id
func (c *MessageController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	if err := c.messageService.DeleteComment(ctx, commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}
