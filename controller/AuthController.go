package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/service"
)

type AuthController struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewAuthController(userService *service.UserService, authService *service.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

// POST /login
func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.UserLoginDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем пользователя по email + password
	user, err := c.userService.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Генерируем access + refresh токены через AuthService
	access, refresh, err := c.authService.GenerateTokens(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate tokens"})
		return
	}

	// Возвращаем JSON с токенами и данными пользователя
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"bio":      user.Bio,
			"avatar":   user.AvatarURL,
		},
	})
}

// POST /refresh
func (c *AuthController) Refresh(ctx *gin.Context) {
	var req dto.TokenRefreshDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccess, newRefresh, err := c.authService.RefreshTokens(ctx, req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}

// POST /logout
func (c *AuthController) Logout(ctx *gin.Context) {
	userID, exists := ctx.Get("userID") // userID получаем из JWT middleware
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := c.authService.Logout(ctx, userID.(string)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
