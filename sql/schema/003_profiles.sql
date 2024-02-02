-- +goose Up
CREATE TABLE IF NOT EXISTS profiles (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    role user_role NOT NULL,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    dob TIMESTAMP NOT NULL,
    bio TEXT,
    avatar VARCHAR(255)
);

-- +goose Down
DROP TABLE profiles;