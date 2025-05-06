package service

import (
	"1337bo4rd/internal/core/domain"
	"1337bo4rd/internal/core/port"
)

type PostService struct {
	repo port.PostRepository
}

func NewPostService(repo port.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(post *domain.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}
	return s.repo.CreatePost(post)
}

func (s *PostService) ListPosts() ([]domain.Post, error) {
	return s.repo.ListPosts()
}

func validatePost(p *domain.Post) error {
	if p.Title == "" {
		return port.ErrEmptyTitle
	} else if p.Content == "" {
		return port.ErrEmptyContent
	} else if p.Avatar == "" {
		return port.ErrEmptyAvatar
	}
	return nil
}
