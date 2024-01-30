-- +goose Up
CREATE TABLE IF NOT EXISTS transaction_members (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    role trans_role NOT NULL
);

-- +goose Down
DROP TABLE transaction_members;