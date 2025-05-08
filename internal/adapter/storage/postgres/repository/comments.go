package repository

import (
	"1337bo4rd/internal/core/domain"
	"database/sql"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) GetLastComment(id *uint64) (*domain.Comment, error) {
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
		return &domain.Comment{}, err
	}

	return &comment, nil
}
