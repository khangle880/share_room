
-- name: CreateEvent :one
INSERT INTO events (name, description, icon_id, background)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEventByID :one
SELECT * FROM events WHERE id = $1;

-- name: GetEventsByIDs :many
SELECT * FROM events WHERE id = ANY($1::UUID[]);