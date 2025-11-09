package service

import (
	"context"
	"time"

	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"
)

type ConversationService struct {
	conversationRepo repository.IConversationRepository
	messageRepo      repository.IDirectMessageRepository
}

func NewConversationService(convRepo repository.IConversationRepository, msgRepo repository.IDirectMessageRepository) *ConversationService {
	return &ConversationService{
		conversationRepo: convRepo,
		messageRepo:      msgRepo,
	}
}

func (s *ConversationService) StartConversation(ctx context.Context, userAID, userBID string) (*model.Conversation, error) {
	existing, err := s.conversationRepo.FindByUsers(ctx, userAID, userBID)
	if err == nil {
		return existing, nil
	}

	newConv := &model.Conversation{
		UserAID:   userAID,
		UserBID:   userBID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.conversationRepo.Create(ctx, newConv); err != nil {
		return nil, err
	}
	return newConv, nil
}

func (s *ConversationService) GetConversations(ctx context.Context, userID string) ([]model.Conversation, error) {
	return s.conversationRepo.GetByUserID(ctx, userID)
}

func (s *ConversationService) GetMessages(ctx context.Context, conversationID string) ([]model.DirectMessage, error) {
	return s.messageRepo.GetByConversationID(ctx, conversationID)
}

func (s *ConversationService) SendMessage(ctx context.Context, msg *model.DirectMessage) error {
	msg.CreatedAt = time.Now()
	return s.messageRepo.Create(ctx, msg)
}
