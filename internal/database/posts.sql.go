// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at,updated_at,title,description,url,feeds_id)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id, created_at, updated_at, title, description, url, feeds_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description sql.NullString
	Url         string
	FeedsID     uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.FeedsID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.FeedsID,
	)
	return i, err
}
