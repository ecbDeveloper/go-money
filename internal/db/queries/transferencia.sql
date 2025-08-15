-- name: CreateTransferencia :exec
INSERT INTO transferencia (
	id_conta, 
	valor,
	tipo
) VALUES ( $1,$2, $3 );