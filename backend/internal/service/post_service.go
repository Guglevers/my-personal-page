package service

import (
	"errors"
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo : repo}
}

func (p *PostService) CreatePost(post domain.Post) error {
    if post.Title == "" || post.Content == "" {
        return errors.New("title and content are required")
    }
    return p.repo.Create(&post)
}

func (p *PostService) GetPosts() ([]domain.Post, error) {
	posts, err := p.repo.Get()
	return posts, err
}