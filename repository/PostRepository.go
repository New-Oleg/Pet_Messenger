package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *PostRepository) Update(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *PostRepository) Delete(ctx context.Context, postID string) error {
	return r.db.WithContext(ctx).Delete(&model.Post{}, "id = ?", postID).Error
}

func (r *PostRepository) GetByID(ctx context.Context, ID string) (*model.Post, error) {
	var post model.Post
	if err := r.db.WithContext(ctx).First(&post, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAllByAuthorID(ctx context.Context, authorID string) ([]model.Post, error) {
	var posts []model.Post
	if err := r.db.WithContext(ctx).Where("author_id = ?", authorID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) AddLike(ctx context.Context, userID string, postID string) error {
	like := &model.Like{
		UserID: userID,
		PostID: postID,
	}

	err := r.db.WithContext(ctx).Create(like).Error
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&model.Post{}).Where("id = ?", postID).UpdateColumn("likes_count", gorm.Expr("likes_count + ?", 1)).Error
}

func (r *PostRepository) RemoveLike(ctx context.Context, userID, postID string) error {
	err := r.db.WithContext(ctx).Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.Like{}).Error
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&model.Post{}).Where("id = ?", postID).UpdateColumn("likes_count", gorm.Expr("likes_count - ?", 1)).Error
}
