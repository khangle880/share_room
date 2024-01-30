-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    type category_type NOT NULL,
    icon_id UUID NOT NULL REFERENCES icons(id),
    parent_id UUID REFERENCES categories(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE categories;