package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IUserRepository interface {
	// Создать нового пользователя
	Create(ctx context.Context, user *model.User) error

	// Найти пользователя по email
	GetByEmail(ctx context.Context, email string) (*model.User, error)

	// Найти пользователя по ID
	GetByID(ctx context.Context, id string) (*model.User, error)

	// Обновить пользователя
	Update(ctx context.Context, user *model.User) error
}
