package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yourname/pet_messenger/config"
	"github.com/yourname/pet_messenger/model"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Like{},
		&model.RefreshToken{},
		&model.DirectMessage{},
		&model.Conversation{},
	)

	log.Println("Database connected successfully!")
	return db
}
