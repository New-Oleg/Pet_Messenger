package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
	"gorm.io/gorm"
)

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) IConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) Create(ctx context.Context, conv *model.Conversation) error {
	return r.db.WithContext(ctx).Create(conv).Error
}

func (r *ConversationRepository) FindByUsers(ctx context.Context, userAID, userBID string) (*model.Conversation, error) {
	var conv model.Conversation
	err := r.db.WithContext(ctx).
		Where("(user_a_id = ? AND user_b_id = ?) OR (user_a_id = ? AND user_b_id = ?)",
			userAID, userBID, userBID, userAID).
		First(&conv).Error
	if err != nil {
		return nil, err
	}
	return &conv, nil
}

func (r *ConversationRepository) GetByUserID(ctx context.Context, userID string) ([]model.Conversation, error) {
	var convs []model.Conversation
	err := r.db.WithContext(ctx).
		Where("user_a_id = ? OR user_b_id = ?", userID, userID).
		Preload("Messages").
		Find(&convs).Error
	return convs, err
}

func (r *ConversationRepository) GetByID(ctx context.Context, id string) (*model.Conversation, error) {
	var conv model.Conversation
	err := r.db.WithContext(ctx).
		Preload("Messages").
		First(&conv, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &conv, nil
}
