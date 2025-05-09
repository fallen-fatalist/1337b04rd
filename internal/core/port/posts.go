package port

import "1337bo4rd/internal/core/domain"

type PostRepository interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
	GetPostWithCommentsById(id *uint64) (*domain.PostComents, error)
}

type PostService interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
	ListActive() ([]domain.Post, error)
	GetPostWithCommentsById(id string) (*domain.PostComents, error)
	CreateComment(comment *domain.Comment, id string) error
}
