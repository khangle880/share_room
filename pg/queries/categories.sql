-- name: CreateCategory :one
INSERT INTO categories (name, type, icon_id, parent_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = $1;

-- name: GetCategoriesByIDs :many
SELECT * FROM categories WHERE id = ANY($1::UUID[]);

-- name: GetCategories :many
SELECT * FROM categories
OFFSET $1
LIMIT $2;

-- name: UpdateCategory :one
UPDATE categories
SET updated_at = NOW(),
    name = COALESCE(sqlc.narg(name), name),
    type = COALESCE(sqlc.narg(type), type),
    icon_id = COALESCE(sqlc.narg(icon_id), icon_id),
    parent_id = COALESCE(sqlc.narg(parent_id), parent_id)
WHERE id = $1
RETURNING *;