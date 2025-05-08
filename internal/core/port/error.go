package port

import "errors"

var (
	ErrEmptyTitle    = errors.New("empty title provided")
	ErrEmptyContent  = errors.New("empty content provided")
	ErrNoPosts       = errors.New("no posts")
	ErrInvalidPostId = errors.New("invalid post ID")
)
