package handler

import (
	"context"
	"encoding/json"
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/service"
	"net/http"
)

type PostHandler struct {
	service *service.PostService
}

type CreatePostResponse struct {
    Message string      `json:"message"`
    Post    domain.Post `json:"post"`
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var input domain.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		http.Error(w, "Request not accepeted", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreatePost(ctx, input);
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreatePostResponse{
    	Message: "Post created successfully",
    	Post:    created,
    })
}

func (h *PostHandler) GetPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) { 
	posts, err := h.service.GetAll(ctx)

	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
        return 
    }
}