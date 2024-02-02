-- name: CreateIcon :one
INSERT INTO icons (name, url, type)
VALUES ($1, $2, $3)
RETURNING *;