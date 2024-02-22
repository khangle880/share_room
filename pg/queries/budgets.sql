-- name: CreateBudget :one
INSERT INTO budgets (name, description, amount, icon_id, room_id, period, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: SetBudgetMember :exec
INSERT INTO budget_members (user_id, budget_id, role)
VALUES ($1, $2, $3);

-- name: UnsetBudgetMembers :exec
DELETE FROM budget_members 
WHERE budget_id = $1;

-- name: GetBudgetByID :one
SELECT * FROM budgets WHERE id = $1;

-- name: GetBudgetsByIDs :many
SELECT * FROM budgets WHERE id = ANY($1::UUID[]);

-- name: GetBudgets :many
SELECT * FROM budgets
OFFSET $1
LIMIT $2;

-- name: GetMembersByBudgetIDs :many
SELECT sqlc.embed(u), bm.budget_id
FROM users u
INNER JOIN budget_members bm ON u.id = bm.user_id
WHERE bm.budget_id = ANY($1::UUID[])
GROUP BY bm.budget_id;

-- name: UpdateBudget :one
UPDATE budgets
SET updated_at = NOW(),
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    amount = COALESCE(sqlc.narg(amount), amount),
    icon_id = COALESCE(sqlc.narg(icon_id), icon_id),
    room_id = COALESCE(sqlc.narg(room_id), room_id),
    period = COALESCE(sqlc.narg(period), period),
    start_date = COALESCE(sqlc.narg(start_date), start_date),
    end_date = COALESCE(sqlc.narg(end_date), end_date)
WHERE id = $1
RETURNING *;

-- name: DeleteBudget :exec
DELETE FROM budgets WHERE id = $1;