package models

import (
	"context"
	"time"

	"github.com/ecbDeveloper/go-money/internal/shared"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateClient struct {
	Categoria      pgtype.Int4     `json:"categoria"`
	Telefone       string          `json:"telefone"`
	Email          string          `json:"email"`
	PessoaFisica   *PessoaFisica   `json:"pessoa_fisica"`
	PessoaJuridica *PessoaJuridica `json:"pessoa_juridica"`
}

type PessoaFisica struct {
	NomeCompleto   string    `json:"nome_completo"`
	DataNascimento time.Time `json:"data_nascimento"`
	Cpf            string    `json:"cpf"`
}

type PessoaJuridica struct {
	DataCriacao  time.Time `json:"data_criacao"`
	NomeFantasia string    `json:"nome_fantasia"`
	Cnpj         string    `json:"cnpj"`
}

func (req CreateClient) Valid(context.Context) shared.ErrorsValidator {
	var eval shared.ErrorsValidator

	eval.CheckField(req.Categoria.Int32 > 0 && req.Categoria.Int32 < 3, "categoria", "esse campo precisa ser uma categoria valida")

	eval.CheckField(shared.IsEmail(req.Email), "email", "esse campo precisa ser um email válido")

	eval.CheckField(shared.NotBlank(req.Telefone), "telefone", "esse campo não pode ser vazio")
	eval.CheckField(shared.NotBlank(req.Email), "email", "esse campo não pode ser vazio")

	eval.CheckField(shared.NotBlank(req.PessoaFisica.Cpf), "cpf", "esse campo não pode ser vazio")
	eval.CheckField(shared.NotBlank(req.PessoaFisica.NomeCompleto), "nome_completo", "esse campo não pode ser vazio")
	eval.CheckField(shared.NotBlank(req.PessoaFisica.DataNascimento.String()), "data_nascimento", "esse campo não pode ser vazio")

	eval.CheckField(shared.NotBlank(req.PessoaJuridica.NomeFantasia), "nome_fantasia", "esse campo não pode ser vazio")
	eval.CheckField(shared.NotBlank(req.PessoaJuridica.Cnpj), "cnpj", "esse campo não pode ser vazio")
	eval.CheckField(shared.NotBlank(req.PessoaJuridica.DataCriacao.String()), "data_criacao", "esse campo não pode ser vazio")

	return eval
}
