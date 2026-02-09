package domain

import (
	"errors"
	"time"
)

type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time
}

var (
	ErrNotFound = errors.New("Post not found")
)
