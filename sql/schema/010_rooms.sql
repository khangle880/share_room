-- +goose Up
CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    avatar VARCHAR(255),
    background VARCHAR(255)
);

-- +goose Down
DROP TABLE rooms;