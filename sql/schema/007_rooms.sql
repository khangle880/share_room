-- +goose Up
CREATE TABLE IF NOT EXISTS rooms (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    avatar VARCHAR(255),
    background VARCHAR(255)
);

-- +goose Down
DROP TABLE rooms;