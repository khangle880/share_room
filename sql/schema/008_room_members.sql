-- +goose Up
CREATE TABLE room_members (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    role room_role NOT NULL
);

-- +goose Down
DROP TABLE room_members;