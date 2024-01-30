-- +goose Up
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    role user_role NOT NULL,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    dob TIMESTAMP NOT NULL,
    bio TEXT,
    avatar VARCHAR(255),
    phone VARCHAR(255)
);

-- +goose Down
DROP TABLE profiles;