-- name: CreateUserProfile :exec
INSERT INTO user_profiles (user_id, profile_id) 
VALUES ($1, $2);