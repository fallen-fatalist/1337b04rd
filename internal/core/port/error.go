package port

import "errors"

var (
	ErrEmptyTitle   = errors.New("empty title provided")
	ErrEmptyContent = errors.New("empty content provided")
	ErrEmptyAvatar  = errors.New("empty Avatar provided")
	ErrNoPosts      = errors.New("no posts")
)
