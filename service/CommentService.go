package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"
)

type CommentService struct {
	repo repository.ICommentRepository
}

func NewCommentService(repo repository.ICommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(ctx context.Context, userID, postID, text string) (*model.Comment, error) {
	if text == "" || userID == "" || postID == "" {
		return nil, errors.New("userID, postID and text are required")
	}

	comment := &model.Comment{
		ID:        uuid.New().String(),
		UserID:    userID,
		PostID:    postID,
		Text:      text,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) GetCommentsByPost(ctx context.Context, postID string) ([]model.Comment, error) {
	return s.repo.GetByPostID(ctx, postID)
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID string) error {
	return s.repo.Delete(ctx, commentID)
}
