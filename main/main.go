package main

// @title Pet Messenger API
// @version 1.0
// @description This is the backend API for the Pet Messenger social app.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.petmessenger.com/support
// @contact.email support@petmessenger.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
import (
	"log"

	"github.com/yourname/pet_messenger/config"
	"github.com/yourname/pet_messenger/router"
)

func main() {
	// Загружаем конфигурацию из .env
	cfg := config.LoadConfig()

	// Инициализация роутера (подключает БД, сервисы, контроллеры)
	r := router.SetupRouter(cfg)

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
