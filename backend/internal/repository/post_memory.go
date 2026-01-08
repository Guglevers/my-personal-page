package repository

import (
	"my-personal-page/backend/internal/domain"
)

type PostRepository interface {
	Create(post *domain.Post) error
}

type PostMemoryRepo struct {
	posts []domain.Post
}

func newPostMemoryRepo() *PostMemoryRepo {
	return &PostMemoryRepo{}
}

func (p *PostMemoryRepo) Create(post domain.Post) {
	p.posts = append(p.posts, post)
}