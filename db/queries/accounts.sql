-- name: CreateAccount :one
INSERT INTO accounts (id, user_id, balance, currency, status, created_at, updated_at, account_type, account_number)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: UpdateAccountBalance :exec
UPDATE accounts SET balance = $2, updated_at = $3 WHERE id = $1;
