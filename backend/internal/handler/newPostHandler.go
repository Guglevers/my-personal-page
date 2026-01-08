package handler

import (
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/service"
	"net/http"
	"encoding/json"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) createPost(w http.ResponseWriter, r http.Request) {
	
	var post domain.Post
	json.NewDecoder(r.Body).Decode(&post)

	h.service.CreatePost(post)
}