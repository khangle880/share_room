-- name: CreateProfile :one
INSERT INTO profiles (created_at, updated_at, role, firstname, lastname, dob, bio, avatar, phone)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsersByIDs :many
SELECT * FROM users WHERE id = ANY($1::UUID[]);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;