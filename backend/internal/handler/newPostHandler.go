package handler

import (
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/service"
	"net/http"
	"encoding/json"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil{
		http.Error(w, "Request not accepeted", http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePost(post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created succssfully"})
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) { 
	posts, err := h.service.GetPosts()

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