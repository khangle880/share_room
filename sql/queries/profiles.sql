-- name: CreateProfile :one
INSERT INTO profiles (role, firstname, lastname, dob, bio, avatar)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProfileByUserID :one
SELECT p.*
FROM profiles p
JOIN user_profiles up ON p.id = up.profile_id
JOIN users u ON up.user_id = u.id
WHERE u.id = $1;

-- name: GetProfilesByUserIDs :many
SELECT p.*
FROM profiles p
JOIN user_profiles up ON p.id = up.profile_id
JOIN users u ON up.user_id = u.id
WHERE u.id = ANY($1::UUID[]);
