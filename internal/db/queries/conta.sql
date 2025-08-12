-- name: CreateAccount :exec
INSERT INTO conta (id_cliente) VALUES ($1);