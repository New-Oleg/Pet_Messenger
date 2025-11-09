package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IConversationRepository interface {
	Create(ctx context.Context, conv *model.Conversation) error
	FindByUsers(ctx context.Context, userAID, userBID string) (*model.Conversation, error)
	GetByUserID(ctx context.Context, userID string) ([]model.Conversation, error)
	GetByID(ctx context.Context, id string) (*model.Conversation, error)
}
