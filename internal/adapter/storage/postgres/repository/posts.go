package repository

import (
	"1337bo4rd/internal/core/domain"
	"database/sql"

	_ "github.com/lib/pq"
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
		if image.Valid {
			post.Image = image.String
		} else {
			post.Image = ""
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, sql.ErrNoRows
	}

	return posts, nil
}

func (r *PostRepository) ListPostComments(id *uint64) (domain.Comment, error) {
	query := `
	SELECT * FROM comments
	WHERE post_id = $1
	ORDER BY comment_id DESC
	LIMIT 1
	`
	rows := r.db.QueryRow(query, id)
	var comment domain.Comment
	var postID sql.NullInt64
	var pcID sql.NullInt64
	if err := rows.Scan(&comment.ID, &comment.UserID, &postID, &pcID, &comment.Content, &comment.CreatedAt); err != nil {
		return domain.Comment{}, err
	}

	if postID.Valid {
		comment.PostID = uint64(postID.Int64)
	} else {
		comment.PostID = 0
	}

	return comment, nil
}
