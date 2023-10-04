-- name: CreatePost :one
INSERT INTO posts (id, created_at,updated_at,title,description,url,feeds_id)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetPost :many
SELECT posts.* FROM posts
JOIN feed_follows ON posts.feeds_id = feed_follows.feeds_id
WHERE feed_follows.user_id = $1
LIMIT $2;