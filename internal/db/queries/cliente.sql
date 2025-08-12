-- name: CreateClient :one
INSERT INTO cliente	(
	categoria_cliente, 
	telefone,
	email
	) VALUES ($1, $2, $3)
RETURNING id;