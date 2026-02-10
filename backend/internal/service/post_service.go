package service

import (
	"context"
	"errors"
	db "my-personal-page/backend/internal/db/generated"
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) Create(ctx context.Context, post domain.Post) (domain.Post, error) {
	if post.Title == "" || post.Content == "" {
		return domain.Post{}, errors.New("title and content are required")
	}
	return p.repo.Create(ctx, &post)
}

func (p *PostService) Get(ctx context.Context, id int64) (domain.Post, error) {
	post, err := p.repo.Get(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

func (p *PostService) GetAll(ctx context.Context) ([]domain.Post, error) {
	posts, err := p.repo.GetAll(ctx)
	if err != nil {
		return []domain.Post{}, err
	}
	return posts, nil
}

func (p *PostService) Delete(ctx context.Context, id int64) (domain.Post, error) {
	post, err := p.repo.Delete(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

func (p *PostService) Update(ctx context.Context, newPost db.UpdatePostsParams) (domain.Post, error) {
	post, err := p.repo.Update(ctx, &newPost)
	if err != nil {
		return domain.Post{}, err
	}
	return post, nil
}
