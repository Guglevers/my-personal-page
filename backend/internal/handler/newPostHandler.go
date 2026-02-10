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

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	post := domain.Post{
		Title: req.Title,
		Content:  req.Content,
	}

	created, err := h.service.Create(ctx, post)
	if err != nil {
		http.Error(w, "failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := h.service.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := h.service.Delete(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	
	var input domain.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	post, err := h.service.Update(ctx, input)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}