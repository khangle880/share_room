-- +goose Up
CREATE TABLE IF NOT EXISTS budgets (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    amount DECIMAL(20, 10) DEFAULT 0,
    icon_id UUID NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
    room_id UUID REFERENCES rooms(id) ON DELETE CASCADE,
    period period_type,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP
);

-- +goose Down
DROP TABLE budgets;