-- +goose Up
CREATE TABLE IF NOT EXISTS icons (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    type VARCHAR(255)
);

-- +goose Down
DROP TABLE icons;