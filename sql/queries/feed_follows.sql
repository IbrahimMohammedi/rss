-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feeds_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;



