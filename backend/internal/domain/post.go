package domain

import "time"

type Post struct {
	ID        int64       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time 
}