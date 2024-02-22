-- name: CreateProfile :one
INSERT INTO profiles (role, firstname, lastname, dob, bio, avatar)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateUserProfile :exec
INSERT INTO user_profiles (user_id, profile_id) 
VALUES ($1, $2);

-- name: GetProfileByUserID :one
SELECT p.*
FROM profiles p
JOIN user_profiles up ON p.id = up.profile_id
JOIN users u ON up.user_id = u.id
WHERE u.id = $1;

-- name: GetProfilesByUserIDs :many
SELECT p.*
FROM profiles p
INNER JOIN user_profiles up ON p.id = up.profile_id
WHERE up.user_id = ANY($1::UUID[]);

-- name: UpdateProfile :one
UPDATE profiles
SET updated_at = NOW(),
    role = COALESCE(sqlc.narg(role), role),
    firstname = COALESCE(sqlc.narg(firstname), firstname),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    dob = COALESCE(sqlc.narg(dob), dob),
    bio = COALESCE(sqlc.narg(bio), bio),
    avatar = COALESCE(sqlc.narg(avatar), avatar)
WHERE id = $1
RETURNING *;