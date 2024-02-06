-- +goose Up
CREATE TABLE events (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    icon_id UUID NOT NULL REFERENCES icons(id),
    background VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE events;