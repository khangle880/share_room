// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: rooms.sql

package pg

import (
	"context"

	"github.com/google/uuid"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (name, address, avatar, background)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, name, address, avatar, background
`

type CreateRoomParams struct {
	Name       string  `json:"name"`
	Address    *string `json:"address"`
	Avatar     *string `json:"avatar"`
	Background *string `json:"background"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error) {
	row := q.db.QueryRow(ctx, createRoom,
		arg.Name,
		arg.Address,
		arg.Avatar,
		arg.Background,
	)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Address,
		&i.Avatar,
		&i.Background,
	)
	return i, err
}

const getMembersByRoomIDs = `-- name: GetMembersByRoomIDs :many
SELECT u.id, u.created_at, u.updated_at, u.deleted_at, u.last_join_at, u.username, u.hashed_password, u.email, u.phone, rm.room_id
FROM users u
INNER JOIN room_members rm ON u.id = rm.user_id
WHERE rm.room_id = ANY($1::UUID[])
`

type GetMembersByRoomIDsRow struct {
	User   User      `json:"user"`
	RoomID uuid.UUID `json:"room_id"`
}

func (q *Queries) GetMembersByRoomIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]GetMembersByRoomIDsRow, error) {
	rows, err := q.db.Query(ctx, getMembersByRoomIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMembersByRoomIDsRow{}
	for rows.Next() {
		var i GetMembersByRoomIDsRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.CreatedAt,
			&i.User.UpdatedAt,
			&i.User.DeletedAt,
			&i.User.LastJoinAt,
			&i.User.Username,
			&i.User.HashedPassword,
			&i.User.Email,
			&i.User.Phone,
			&i.RoomID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoomByID = `-- name: GetRoomByID :one
SELECT id, created_at, updated_at, name, address, avatar, background FROM rooms WHERE id = $1
`

func (q *Queries) GetRoomByID(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoomByID, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Address,
		&i.Avatar,
		&i.Background,
	)
	return i, err
}

const getRooms = `-- name: GetRooms :many
SELECT id, created_at, updated_at, name, address, avatar, background FROM rooms
OFFSET $1
LIMIT $2
`

type GetRoomsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetRooms(ctx context.Context, arg GetRoomsParams) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRooms, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Address,
			&i.Avatar,
			&i.Background,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoomsByIDs = `-- name: GetRoomsByIDs :many
SELECT id, created_at, updated_at, name, address, avatar, background FROM rooms WHERE id = ANY($1::UUID[])
`

func (q *Queries) GetRoomsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRoomsByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Address,
			&i.Avatar,
			&i.Background,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByRoomRole = `-- name: GetUserByRoomRole :one
SELECT u.id, u.created_at, u.updated_at, u.deleted_at, u.last_join_at, u.username, u.hashed_password, u.email, u.phone
FROM users u
INNER JOIN room_members rm ON u.id = rm.user_id
WHERE rm.role = $1 AND rm.room_id = $2
`

type GetUserByRoomRoleParams struct {
	Role   RoomRole  `json:"role"`
	RoomID uuid.UUID `json:"room_id"`
}

func (q *Queries) GetUserByRoomRole(ctx context.Context, arg GetUserByRoomRoleParams) (User, error) {
	row := q.db.QueryRow(ctx, getUserByRoomRole, arg.Role, arg.RoomID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.LastJoinAt,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Phone,
	)
	return i, err
}

const setRoomMember = `-- name: SetRoomMember :exec
INSERT INTO room_members (user_id, room_id, role)
VALUES ($1, $2, $3)
`

type SetRoomMemberParams struct {
	UserID uuid.UUID `json:"user_id"`
	RoomID uuid.UUID `json:"room_id"`
	Role   RoomRole  `json:"role"`
}

func (q *Queries) SetRoomMember(ctx context.Context, arg SetRoomMemberParams) error {
	_, err := q.db.Exec(ctx, setRoomMember, arg.UserID, arg.RoomID, arg.Role)
	return err
}

const unsetRoomMembers = `-- name: UnsetRoomMembers :exec
DELETE FROM room_members 
WHERE room_id = $1 AND role = room_role.member
`

func (q *Queries) UnsetRoomMembers(ctx context.Context, roomID uuid.UUID) error {
	_, err := q.db.Exec(ctx, unsetRoomMembers, roomID)
	return err
}

const updateRoom = `-- name: UpdateRoom :one
UPDATE rooms
SET updated_at = NOW(),
    name = COALESCE($2, name),
    address = COALESCE($3, address),
    avatar = COALESCE($4, avatar),
    background = COALESCE($5, background)
WHERE id = $1
RETURNING id, created_at, updated_at, name, address, avatar, background
`

type UpdateRoomParams struct {
	ID         uuid.UUID `json:"id"`
	Name       *string   `json:"name"`
	Address    *string   `json:"address"`
	Avatar     *string   `json:"avatar"`
	Background *string   `json:"background"`
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error) {
	row := q.db.QueryRow(ctx, updateRoom,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.Avatar,
		arg.Background,
	)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Address,
		&i.Avatar,
		&i.Background,
	)
	return i, err
}