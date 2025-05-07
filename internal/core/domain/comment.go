package domain

import "time"

type Comment struct {
	ID              uint64    `json:"comment_id"`
	UserID          uint64    `json:"user_id"`
	PostID          uint64    `json:"post_id"`
	ParentCommentID uint64    `json:"parent_comment_id,omitempty"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
}
