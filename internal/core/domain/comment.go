package domain

import "time"

type Comment struct {
	ID              uint64
	UserID          string
	PostID          uint64
	ParentCommentID uint64
	Content         string
	CreatedAt       time.Time
}
