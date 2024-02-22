-- name: CreateIcon :one
INSERT INTO icons (name, url, type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetIconByID :one
SELECT * FROM icons WHERE id = $1;

-- name: GetIconsByIDs :many
SELECT * FROM icons WHERE id = ANY($1::UUID[]);

