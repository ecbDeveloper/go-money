-- name: CreateAccount :one
INSERT INTO conta (id_cliente) VALUES ($1)
RETURNING id;

-- name: GetAllAccountsByClientId :many
SELECT * FROM conta
WHERE id_cliente = $1;

-- name: GetBalanceByAccountId :one
SELECT saldo FROM conta
WHERE id = $1;