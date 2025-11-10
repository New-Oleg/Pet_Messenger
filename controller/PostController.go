package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/service"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{postService: postService}
}

// POST /posts
// @Summary Create a post
// @Description Create a new post for the authenticated user
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dto.PostCreateDTO true "Post data"
// @Success 201 {object} dto.PostResponse
// @Failure 400 {object} map[string]string
// @Router /posts [post]
func (c *PostController) CreatePost(ctx *gin.Context) {
	var req dto.PostCreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID := ctx.GetString("userID") // JWT middleware
	post, err := c.postService.CreatePost(ctx, authorID, req.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, toPostResponse(post))
}

// GET /posts/:id
// @Summary Get post by ID
// @Description Get post details
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} dto.PostResponse
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [get]
func (c *PostController) GetPostByID(ctx *gin.Context) {
	postID := ctx.Param("id")
	post, err := c.postService.GetPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	ctx.JSON(http.StatusOK, toPostResponse(post))
}

// GET /users/:id/posts
// @Summary Get posts by user
// @Description Get all posts created by a specific user
// @Tags posts
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} dto.PostResponse
// @Router /users/{id}/posts [get]
func (c *PostController) GetPostsByAuthor(ctx *gin.Context) {
	authorID := ctx.Param("id")
	posts, err := c.postService.GetPostsByAuthor(ctx, authorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.PostResponse, len(posts))
	for i, p := range posts {
		responses[i] = toPostResponse(&p)
	}

	ctx.JSON(http.StatusOK, responses)
}

// PUT /posts/:id
// @Summary Update post
// @Description Update a post by the authenticated user
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body dto.PostCreateDTO true "Post data"
// @Success 200 {object} dto.PostResponse
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [put]
func (c *PostController) UpdatePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	var req dto.PostCreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := c.postService.GetPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	if post.AuthorID != ctx.GetString("userID") {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you can update only your posts"})
		return
	}

	post.Text = req.Text
	if err := c.postService.UpdatePost(ctx, post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, toPostResponse(post))
}

// DELETE /posts/:id
// @Summary Delete post
// @Description Delete a post by the authenticated user
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [delete]
func (c *PostController) DeletePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	post, err := c.postService.GetPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	if post.AuthorID != ctx.GetString("userID") {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you can delete only your posts"})
		return
	}

	if err := c.postService.DeletePost(ctx, postID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

// POST /posts/:id/like
// @Summary Like post
// @Description Like a post by the authenticated user
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Router /posts/{id}/like [post]
func (c *PostController) LikePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetString("userID")

	if err := c.postService.LikePost(ctx, userID, postID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post liked"})
}

// POST /posts/:id/unlike
// @Summary Unlike post
// @Description Remove like from a post by the authenticated user
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Router /posts/{id}/unlike [post]
func (c *PostController) UnlikePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetString("userID")

	if err := c.postService.UnlikePost(ctx, userID, postID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "like removed"})
}

// --- Вспомогательная функция, потому что я поленился сделать маппер))))
func toPostResponse(post *model.Post) dto.PostResponse {
	return dto.PostResponse{
		ID:         post.ID,
		AuthorID:   post.AuthorID,
		Text:       post.Text,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
		LikesCount: post.LikesCount,
	}
}
