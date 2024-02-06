-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    budget_id UUID REFERENCES budgets(id) ON DELETE CASCADE,
    event_id UUID REFERENCES events(id) ON DELETE CASCADE,
    exc_time TIMESTAMP NOT NULL,
    description TEXT,
    amount DECIMAL(20, 10) DEFAULT 0,
    images VARCHAR (255) []
);

-- +goose Down
DROP TABLE transactions;