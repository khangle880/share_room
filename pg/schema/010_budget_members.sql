-- +goose Up
CREATE TABLE IF NOT EXISTS budget_members (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    budget_id UUID NOT NULL REFERENCES budgets(id) ON DELETE CASCADE,
    role budget_role NOT NULL
);

-- +goose Down
DROP TABLE budget_members;