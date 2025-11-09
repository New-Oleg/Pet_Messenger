package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourname/pet_messenger/config"
	"github.com/yourname/pet_messenger/controller"
	"github.com/yourname/pet_messenger/dto"
	"github.com/yourname/pet_messenger/middleware"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/pkg/db"
	"github.com/yourname/pet_messenger/repository"
	"github.com/yourname/pet_messenger/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Подключаем базу
	gormDB := db.ConnectDB(cfg)

	// Миграции
	gormDB.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Like{},
		&model.RefreshToken{},
		&model.Conversation{},
		&model.DirectMessage{},
	)

	// --- Репозитории ---
	userRepo := repository.NewUserRepository(gormDB)
	postRepo := repository.NewPostRepository(gormDB)
	commentRepo := repository.NewCommentRepository(gormDB)
	refreshRepo := repository.NewRefreshTokenRepository(gormDB)
	conversationRepo := repository.NewConversationRepository(gormDB)
	messageRepo := repository.NewDirectMessageRepository(gormDB)

	// --- Сервисы ---
	userService := service.NewUserService(userRepo, cfg.JWTSecret, 15*time.Minute)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo)
	authService := service.NewAuthService(cfg.JWTSecret, 15*time.Minute, 7*24*time.Hour, refreshRepo)
	conversationService := service.NewConversationService(conversationRepo, messageRepo)

	// --- Контроллеры ---
	authCtrl := controller.NewAuthController(userService, authService)
	postCtrl := controller.NewPostController(postService)
	commentCtrl := controller.NewCommentController(commentService)
	conversationCtrl := controller.NewConversationController(conversationService)

	// --- Публичные маршруты ---
	r.POST("/register", func(ctx *gin.Context) {
		var req dto.UserRegisterDTO
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.Register(ctx, req.Username, req.Email, req.Password)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"bio":      user.Bio,
			"avatar":   user.AvatarURL,
		})
	})
	r.POST("/login", authCtrl.Login)
	r.POST("/refresh", authCtrl.Refresh)

	// --- Защищённые маршруты ---
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	auth.POST("/logout", authCtrl.Logout)

	auth.GET("/profile", func(ctx *gin.Context) {
		userID := ctx.GetString("userID")
		user, err := userService.GetProfile(ctx, userID)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(200, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"bio":      user.Bio,
			"avatar":   user.AvatarURL,
		})
	})

	// --- Посты ---
	auth.POST("/posts", postCtrl.CreatePost)
	auth.GET("/posts/:id", postCtrl.GetPostByID)
	auth.GET("/users/:id/posts", postCtrl.GetPostsByAuthor)
	auth.PUT("/posts/:id", postCtrl.UpdatePost)
	auth.DELETE("/posts/:id", postCtrl.DeletePost)
	auth.POST("/posts/:id/like", postCtrl.LikePost)
	auth.POST("/posts/:id/unlike", postCtrl.UnlikePost)

	// --- Комментарии ---
	auth.POST("/posts/:id/comments", commentCtrl.CreateComment)
	auth.GET("/posts/:id/comments", commentCtrl.GetCommentsByPost)
	auth.DELETE("/comments/:id", commentCtrl.DeleteComment)

	// ---
	auth.POST("/conversations", conversationCtrl.StartConversation)
	auth.GET("/conversations", conversationCtrl.GetConversations)
	auth.GET("/conversations/:id/messages", conversationCtrl.GetMessages)
	auth.POST("/conversations/:id/messages", conversationCtrl.SendMessage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
