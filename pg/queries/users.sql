-- name: CreateUser :one
INSERT INTO users (last_join_at, username, hashed_password, email, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUsersByIDs :many
SELECT * FROM users WHERE id = ANY($1::UUID[]) AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 AND deleted_at IS NULL;

-- name: UpdateUser :one
UPDATE users
SET updated_at = NOW(),
    last_join_at = COALESCE(sqlc.narg(last_join_at), last_join_at),
    username = COALESCE(sqlc.narg(username), username),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    email = COALESCE(sqlc.narg(email), email),
    phone = COALESCE(sqlc.narg(phone), phone)
WHERE id = $1 AND deleted_at IS NULL 
RETURNING *;

-- name: DeleteUser :exec
UPDATE users SET deleted_at = NOW() WHERE id = $1;
