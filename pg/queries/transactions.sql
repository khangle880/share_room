-- name: CreateTransaction :one
INSERT INTO transactions (category_id, budget_id, event_id, exc_time, description, amount, images)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: SetTranMember :exec
INSERT INTO transaction_members (user_id, transaction_id, role)
VALUES ($1, $2, $3);

-- name: UnsetTranMembers :exec
DELETE FROM transaction_members 
WHERE transaction_id = $1;

-- name: GetTransactionByID :one
SELECT * FROM transactions WHERE id = $1;

-- name: GetTransactionsByIDs :many
SELECT * FROM transactions WHERE id = ANY($1::UUID[]);

-- name: GetTransByBudgetIDs :many
SELECT * FROM transactions WHERE budget_id = ANY($1::UUID[]);

-- name: GetTransactions :many
SELECT * FROM transactions
OFFSET $1
LIMIT $2;

-- name: GetMembersByTranIDs :many
SELECT sqlc.embed(u), tm.transaction_id
FROM users u
INNER JOIN transaction_members tm ON u.id = tm.user_id
WHERE tm.room_id = ANY($1::UUID[]) AND tm.role = $2;

-- name: UpdateTransaction :one
UPDATE transactions
SET updated_at = NOW(),
    category_id = COALESCE(sqlc.narg(category_id), category_id),
    budget_id = COALESCE(sqlc.narg(budget_id), budget_id),
    event_id = COALESCE(sqlc.narg(event_id), event_id),
    exc_time = COALESCE(sqlc.narg(exc_time), exc_time),
    description = COALESCE(sqlc.narg(description), description),
    amount = COALESCE(sqlc.narg(amount), amount),
    images = COALESCE(sqlc.narg(images), images)
WHERE id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1;