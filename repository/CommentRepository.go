package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, message *model.Comment) error {
	return r.db.WithContext(ctx).Create(message).Error
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID string) ([]model.Comment, error) {
	var messages []model.Comment
	if err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *CommentRepository) Delete(ctx context.Context, messageID string) error {
	return r.db.WithContext(ctx).Delete(&model.Comment{}, "id = ?", messageID).Error
}
