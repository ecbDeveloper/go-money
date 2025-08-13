-- name: CreateAccount :one
INSERT INTO conta (id_cliente) VALUES ($1)
RETURNING id;