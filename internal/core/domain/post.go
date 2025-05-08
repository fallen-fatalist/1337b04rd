package domain

import "time"

type Post struct {
	ID        uint64
	UserID    string
	Title     string
	Content   string
	Image     string
	CreatedAt time.Time
}

type PostComents struct {
	Post     Post
	Comments []Comment
}
