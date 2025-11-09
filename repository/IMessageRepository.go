package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IMessageRepository interface {
	// Создать комментарий к посту
	Create(ctx context.Context, message *model.Message) error

	// Получить все комментарии к посту
	GetByPostID(ctx context.Context, postID string) ([]model.Message, error)

	// Удалить комментарий по ID
	Delete(ctx context.Context, messageID string) error
}
