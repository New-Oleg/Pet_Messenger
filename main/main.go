package main

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
