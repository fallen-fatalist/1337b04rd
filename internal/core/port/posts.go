package port

import "1337bo4rd/internal/core/domain"

type PostRepository interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
	ListPostComments(id *uint64) (domain.Comment, error)
}

type PostService interface {
	CreatePost(post *domain.Post) error
	ListPosts() ([]domain.Post, error)
	ListActive() ([]domain.Post, error)
}
