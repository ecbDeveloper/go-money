package models

import (
	"time"

	"github.com/ecbDeveloper/go-money/internal/shared"
)

type CreateClient struct {
	Categoria      int32           `json:"categoria"`
	Telefone       string          `json:"telefone"`
	Email          string          `json:"email"`
	Senha          string          `json:"senha"`
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

func (req CreateClient) Validate() map[string]string {
	errors := make(map[string]string)

	if req.Categoria < 1 || req.Categoria > 2 {
		errors["categoria"] = "esse campo precisa ser uma categoria valida"
	}

	if req.Categoria == 1 && req.PessoaFisica == nil {
		errors["pessoa_fisica"] = "é nescessário preencher os dados de pessoa física pra continuar"
	}

	if req.Categoria == 2 && req.PessoaJuridica == nil {
		errors["pessoa_juridica"] = "é nescessário preencher os dados de pessoa jurídica pra continuar"
	}

	if req.PessoaFisica == nil && req.PessoaJuridica == nil {
		errors["dados_pessoais"] = "é nescessário preencher os dados de pessoa física ou jurídica pra continuar"
	}

	if !shared.NotBlank(req.Telefone) {
		errors["telefone"] = "esse campo não pode ser vazio"
	}

	if !shared.NotBlank(req.Email) {
		errors["email"] = "esse campo não pode ser vazio"
	} else if !shared.IsEmail(req.Email) {
		errors["email"] = "esse campo precisa ser um email válido"
	}

	if !shared.NotBlank(req.Senha) {
		errors["senha"] = "esse campo não pode ser vazio"
	} else if len(req.Senha) < 8 || len(req.Senha) > 35 {
		errors["senha"] = "esse campo precisa ter o tamnho entre 8 e 35 caracteres"
	}

	if req.PessoaFisica != nil {
		if !shared.NotBlank(req.PessoaFisica.Cpf) {
			errors["cpf"] = "esse campo não pode ser vazio"
		}

		if !shared.NotBlank(req.PessoaFisica.NomeCompleto) {
			errors["nome_completo"] = "esse campo não pode ser vazio"
		}

		if req.PessoaFisica.DataNascimento.IsZero() {
			errors["data_nascimento"] = "esse campo não pode ser vazio"
		}
	}

	if req.PessoaJuridica != nil {
		if !shared.NotBlank(req.PessoaJuridica.NomeFantasia) {
			errors["nome_fantasia"] = "esse campo não pode ser vazio"
		}

		if !shared.NotBlank(req.PessoaJuridica.Cnpj) {
			errors["cnpj"] = "esse campo não pode ser vazio"
		}

		if req.PessoaJuridica.DataCriacao.IsZero() {
			errors["data_criacao"] = "esse campo não pode ser vazio"
		}
	}

	return errors
}

type AuthenticateClient struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (req AuthenticateClient) Validate() map[string]string {
	errors := make(map[string]string)

	if !shared.NotBlank(req.Email) {
		errors["email"] = "esse campo precisa ser preenchido"
	} else if !shared.IsEmail(req.Email) {
		errors["email"] = "esse campo precisa ser um email válido"
	}

	if !shared.NotBlank(req.Senha) {
		errors["senha"] = "esse campo precisa ser preenchido"
	}

	return errors
}
