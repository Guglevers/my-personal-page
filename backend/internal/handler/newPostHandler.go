package handler

import (
	"encoding/json"
	"errors"
	"my-personal-page/backend/internal/domain"
	"my-personal-page/backend/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreatePost(ctx, input)
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

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := h.service.Get(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := h.service.Delete(ctx, id)

	if err != nil {
		if
	}

}
