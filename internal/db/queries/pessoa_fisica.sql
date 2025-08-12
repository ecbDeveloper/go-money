-- name: CreatePessoaFisica :exec
INSERT INTO pessoa_fisica(
	id_cliente,
	nome_completo,
	data_nascimento,
	cpf
) VALUES ($1, $2, $3, $4);