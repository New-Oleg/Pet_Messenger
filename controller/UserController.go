package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{service: s}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user account with username, email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterDTO true "User registration data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req dto.UserRegisterDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем DTO
	resp := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}

	ctx.JSON(http.StatusOK, resp)
}
