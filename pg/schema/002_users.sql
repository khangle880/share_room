-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    last_join_at TIMESTAMP NOT NULL,
    username TEXT UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(255),
    CONSTRAINT email_or_phone_not_null CHECK (email IS NOT NULL OR phone IS NOT NULL),
    UNIQUE(username, email, phone, deleted_at)
);

-- CREATE VIEW users AS
-- SELECT *
-- FROM accounts
-- WHERE deleted_at IS NULL;

-- +goose Down
DROP TABLE users;