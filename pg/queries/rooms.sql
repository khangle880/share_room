-- name: CreateRoom :one
INSERT INTO rooms (name, address, avatar, background)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: SetRoomMember :exec
INSERT INTO room_members (user_id, room_id, role)
VALUES ($1, $2, $3);

-- name: UnsetRoomMembers :exec
DELETE FROM room_members 
WHERE room_id = $1 AND role = room_role.member;

-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = $1;

-- name: GetUserByRoomRole :one
SELECT u.*
FROM users u
INNER JOIN room_members rm ON u.id = rm.user_id
WHERE rm.role = $1 AND rm.room_id = $2;

-- name: GetRoomsByIDs :many
SELECT * FROM rooms WHERE id = ANY($1::UUID[]);

-- name: GetRooms :many
SELECT * FROM rooms
OFFSET $1
LIMIT $2;

-- name: GetMembersByRoomIDs :many
SELECT sqlc.embed(u), rm.room_id
FROM users u
INNER JOIN room_members rm ON u.id = rm.user_id
WHERE rm.room_id = ANY($1::UUID[]);

-- name: UpdateRoom :one
UPDATE rooms
SET updated_at = NOW(),
    name = COALESCE(sqlc.narg(name), name),
    address = COALESCE(sqlc.narg(address), address),
    avatar = COALESCE(sqlc.narg(avatar), avatar),
    background = COALESCE(sqlc.narg(background), background)
WHERE id = $1
RETURNING *;