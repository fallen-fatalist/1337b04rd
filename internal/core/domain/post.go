package domain

import "time"

type Post struct {
	ID        uint64    `json:"post_id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
