package service

import (
	"1337bo4rd/internal/core/domain"
	"1337bo4rd/internal/core/port"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type PostService struct {
	repo       port.PostRepository
	commentSvc *CommentService
}

func NewPostService(repo port.PostRepository, commentSvc *CommentService) *PostService {
	return &PostService{
		repo:       repo,
		commentSvc: commentSvc,
	}
}

func (s *PostService) CreatePost(post *domain.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}
	return s.repo.CreatePost(post)
}

func (s *PostService) ListPosts() ([]domain.Post, error) {
	posts, err := s.repo.ListPosts()
	if err != nil {
		return nil, err
	}
	// add logic for deleted threads

	return posts, nil

}

func (s *PostService) ListActive() ([]domain.Post, error) {
	posts, err := s.repo.ListPosts()
	if err != nil {
		return nil, err
	}

	var validPosts []domain.Post
	now := time.Now()

	for _, post := range posts {
		comment, err := s.commentSvc.repo.GetLastComment(&post.ID)
		if err != nil {
			// No comment found, fallback to post time
			if errors.Is(err, sql.ErrNoRows) {
				if now.Sub(post.CreatedAt) < 10*time.Minute {
					validPosts = append(validPosts, post)
				}
				continue
			}
			// Some other error
			return nil, err
		}

		// If comment exists, check if it's recent enough
		if now.Sub(comment.CreatedAt) < 15*time.Minute {
			validPosts = append(validPosts, post)
		}
	}

	return validPosts, nil
}

func (s *PostService) GetPostById(idStr string) (*domain.Post, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	uid := uint64(id)
	post, err := s.repo.GetPostById(&uid)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func validatePost(p *domain.Post) error {
	if p.Title == "" {
		return port.ErrEmptyTitle
	} else if p.Content == "" {
		return port.ErrEmptyContent
	}
	return nil
}
