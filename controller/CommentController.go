package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/service"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetString("userID")

	var req dto.CommentCreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := c.commentService.CreateComment(ctx, userID, postID, req.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// GET /posts/:id/comments
func (c *CommentController) GetCommentsByPost(ctx *gin.Context) {
	postID := ctx.Param("id")

	comments, err := c.commentService.GetCommentsByPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// DELETE /comments/:id
func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	if err := c.commentService.DeleteComment(ctx, commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}
