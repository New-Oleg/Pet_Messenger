package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/service"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

// POST /posts/:id/comments
// @Summary Create comment
// @Description Add a comment to a post
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param comment body dto.CommentCreateDTO true "Comment data"
// @Success 201 {object} dto.CommentResponse
// @Router /posts/{id}/comments [post]
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

	ctx.JSON(http.StatusCreated, toCommentResponse(comment))
}

// GET /posts/:id/comments
// @Summary Get comments
// @Description Get all comments for a post
// @Tags comments
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {array} dto.CommentResponse
// @Router /posts/{id}/comments [get]
func (c *CommentController) GetCommentsByPost(ctx *gin.Context) {
	postID := ctx.Param("id")

	comments, err := c.commentService.GetCommentsByPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.CommentResponse, len(comments))
	for i, cm := range comments {
		responses[i] = toCommentResponse(&cm)
	}

	ctx.JSON(http.StatusOK, responses)
}

// DELETE /comments/:id
// @Summary Delete comment
// @Description Delete a comment by ID
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} map[string]string
// @Router /comments/{id} [delete]
func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	if err := c.commentService.DeleteComment(ctx, commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}

// --- Вспомогательная функция ---
func toCommentResponse(comment *model.Comment) dto.CommentResponse {
	return dto.CommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PostID:    comment.PostID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}
