package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(ctx context.Context, message *model.Message) error {
	return r.db.WithContext(ctx).Create(message).Error
}

func (r *MessageRepository) GetByPostID(ctx context.Context, postID string) ([]model.Message, error) {
	var messages []model.Message
	if err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageRepository) Delete(ctx context.Context, messageID string) error {
	return r.db.WithContext(ctx).Delete(&model.Message{}, "id = ?", messageID).Error
}
