package repository

import (
	"my-personal-page/backend/internal/domain"
)

type PostRepository interface {
	Create(post *domain.Post) error
	Get() ([]domain.Post, error)
}

type PostMemoryRepo struct {
	posts []domain.Post
}

func NewPostMemoryRepo() *PostMemoryRepo{
	return &PostMemoryRepo{
		posts: make([]domain.Post, 0),
	}
}

func (p *PostMemoryRepo) Create(post *domain.Post) error{
	p.posts = append(p.posts, *post)
	return nil  
}

func (p *PostMemoryRepo) Get() ([]domain.Post, error) {
	return p.posts, nil
}