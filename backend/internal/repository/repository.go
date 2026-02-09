package repository

import (
	"context"
	"my-personal-page/backend/internal/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) (domain.Post, error)
	GetAll(ctx context.Context) ([]domain.Post, error)
	Get(ctx context.Context, id int64) (domain.Post, error)
	Delete(ctx context.Context, id int64) (domain.Post, error)
	Update(ctx context.Context, post *domain.Post) (domain.Post, error)
}