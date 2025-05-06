package port

import "1337bo4rd/internal/core/domain"

type PostRepository interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
}

type PostService interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
}
