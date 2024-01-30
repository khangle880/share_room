-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, last_join_at, username, hashed_password, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsersByIDs :many
SELECT * FROM users WHERE id = ANY($1::UUID[]);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;