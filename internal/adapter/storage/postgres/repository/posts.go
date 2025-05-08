package repository

import (
	"1337bo4rd/internal/core/domain"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(post *domain.Post) error {
	var (
		query string
		args  []interface{}
	)
	if post.Image != "" {
		query = `
		INSERT INTO posts (user_id, title, content, image)
		VALUES ($1, $2, $3, $4)
		`
		args = []interface{}{post.UserID, post.Title, post.Content, post.Image}
	} else {
		query = `
		INSERT INTO posts (user_id, title, content)
		VALUES ($1, $2, $3)
		`
		args = []interface{}{post.UserID, post.Title, post.Content}
	}

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) ListPosts() ([]domain.Post, error) {
	var posts []domain.Post
	query := `
	SELECT * FROM posts;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		var image sql.NullString
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &image, &post.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, sql.ErrNoRows
	}

	return posts, nil
}

func (r *PostRepository) GetPostById(id *uint64) (*domain.Post, error) {
	query := `
	SELECT * FROM posts
	WHERE post_id = $1
	`
	var post domain.Post
	var image sql.NullString
	if err := r.db.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &image, &post.CreatedAt); err != nil {
		return nil, err
	}

	return &post, nil
}
