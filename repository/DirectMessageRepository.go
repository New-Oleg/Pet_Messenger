package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
	"gorm.io/gorm"
)

type DirectMessageRepository struct {
	db *gorm.DB
}

func NewDirectMessageRepository(db *gorm.DB) IDirectMessageRepository {
	return &DirectMessageRepository{db: db}
}

func (r *DirectMessageRepository) Create(ctx context.Context, msg *model.DirectMessage) error {
	return r.db.WithContext(ctx).Create(msg).Error
}

func (r *DirectMessageRepository) GetByConversationID(ctx context.Context, convID string) ([]model.DirectMessage, error) {
	var msgs []model.DirectMessage
	err := r.db.WithContext(ctx).Where("conversation_id = ?", convID).Find(&msgs).Error
	return msgs, err
}
