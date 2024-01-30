-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    budget_id UUID REFERENCES budgets(id) ON DELETE CASCADE,
    event_id UUID REFERENCES events(id) ON DELETE CASCADE,
    exc_time TIMESTAMP NOT NULL,
    description TEXT,
    amount DECIMAL(12, 2) DEFAULT 0,
    images VARCHAR (255) []
);

-- +goose Down
DROP TABLE transactions;