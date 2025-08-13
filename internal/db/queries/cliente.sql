-- name: CreateClient :one
INSERT INTO cliente	(
	categoria_cliente, 
	telefone,
	email,
	password
	) VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetClientByEmail :one
SELECT * FROM cliente
WHERE email = $1;