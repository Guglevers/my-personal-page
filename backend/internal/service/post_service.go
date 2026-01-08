package service

import (
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func newPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo : repo}
}

func (p *PostService) CreatePost(post domain.Post) {
	p.repo.Create(&post)
}