package repository

import (
	"context"

	"github.com/yourname/pet_messenger/model"
)

type IPostRepository interface {
	Create(ctx context.Context, post *model.Post) error

	Update(ctx context.Context, post *model.Post) error

	Delete(ctx context.Context, postId string) error

	GetByID(ctx context.Context, ID string) (*model.Post, error)

	GetAllByAuthorID(ctx context.Context, ID string) ([]model.Post, error)

	AddLike(ctx context.Context, userID string, postID string) error

	RemoveLike(ctx context.Context, userID, postID string) error
}
