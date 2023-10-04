-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL UNIQUE,
    feeds_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose Down
DROP TABLE posts;