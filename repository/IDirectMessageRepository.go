package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IDirectMessageRepository interface {
	Create(ctx context.Context, msg *model.DirectMessage) error
	GetByConversationID(ctx context.Context, convID string) ([]model.DirectMessage, error)
}
