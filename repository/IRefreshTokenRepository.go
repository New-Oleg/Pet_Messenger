package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IRefreshTokenRepository interface {
	Create(ctx context.Context, token *model.RefreshToken) error
	GetByToken(ctx context.Context, tokenString string) (*model.RefreshToken, error)
	Delete(ctx context.Context, tokenString string) error
	DeleteAllByUserID(ctx context.Context, userID string) error
}
