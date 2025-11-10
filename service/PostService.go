package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"
)

type PostService struct {
	repo repository.IPostRepository
}

func NewPostService(repo repository.IPostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, authorID, text string) (*model.Post, error) {
	if text == "" {
		return nil, errors.New("text cannot be empty")
	}

	post := &model.Post{
		ID:         uuid.New().String(),
		AuthorID:   authorID,
		Text:       text,
		LikesCount: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

// Получить пост по ID
func (s *PostService) GetPost(ctx context.Context, postID string) (*model.Post, error) {
	return s.repo.GetByID(ctx, postID)
}

// Получить все посты пользователя
func (s *PostService) GetPostsByAuthor(ctx context.Context, authorID string) ([]model.Post, error) {
	return s.repo.GetAllByAuthorID(ctx, authorID)
}

// Обновление поста
func (s *PostService) UpdatePost(ctx context.Context, post *model.Post) error {
	post.UpdatedAt = time.Now()
	return s.repo.Update(ctx, post)
}

// Удаление поста
func (s *PostService) DeletePost(ctx context.Context, postID string) error {
	return s.repo.Delete(ctx, postID)
}

// Поставить лайк
func (s *PostService) LikePost(ctx context.Context, userID, postID string) error {
	return s.repo.AddLike(ctx, userID, postID)
}

// Снять лайк
func (s *PostService) UnlikePost(ctx context.Context, userID, postID string) error {
	return s.repo.RemoveLike(ctx, userID, postID)
}
