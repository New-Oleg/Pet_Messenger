package service

import (
	"context"
	"errors"
	"time"

	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"
)

type MessageService struct {
	repo repository.IMessageRepository
}

// Конструктор
func NewMessageService(repo repository.IMessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// Создать комментарий к посту
func (s *MessageService) CreateComment(ctx context.Context, userID, postID string) (*model.Message, error) {
	if userID == "" || postID == "" {
		return nil, errors.New("userID and postID are required")
	}

	msg := &model.Message{
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

// Получить все комментарии под постом
func (s *MessageService) GetCommentsByPost(ctx context.Context, postID string) ([]model.Message, error) {
	if postID == "" {
		return nil, errors.New("postID is required")
	}
	return s.repo.GetByPostID(ctx, postID)
}

// Удалить комментарий
func (s *MessageService) DeleteComment(ctx context.Context, messageID string) error {
	if messageID == "" {
		return errors.New("messageID is required")
	}
	return s.repo.Delete(ctx, messageID)
}
