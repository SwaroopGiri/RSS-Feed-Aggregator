-- name: CreateFeed :one
INSERT INTO feeds (id, name, created_at, updated_at, url, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeeds :many
SELECT * FROM feeds ORDER BY last_fetched ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds SET last_fetched = NOW(), updated_at = NOW() WHERE id = $1 RETURNING *;