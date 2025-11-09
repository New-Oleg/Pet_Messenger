package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type ICommentRepository interface {
	Create(ctx context.Context, comment *model.Comment) error
	GetByPostID(ctx context.Context, postID string) ([]model.Comment, error)
	Delete(ctx context.Context, commentID string) error
}
