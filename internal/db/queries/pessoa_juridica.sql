-- name: CreatePessoaJuridica :exec
INSERT INTO pessoa_juridica(
	id_cliente,
	data_criacao,
	nome_fantasia,
	cnpj
) VALUES ($1, $2, $3, $4);