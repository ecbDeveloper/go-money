package models

import "github.com/ecbDeveloper/go-money/internal/shared"

type AccountTransaction struct {
	Valor        float64 `json:"valor"`
	IdConta      string  `json:"id_conta"`
	TipoOperacao int32   `json:"tipo_operacao"`
}

func (d AccountTransaction) Validate() map[string]string {
	errors := make(map[string]string)

	if d.Valor < 0.01 {
		errors["valor"] = "depósito precisa ser maior que R$0,00"
	}

	if shared.IsBlank(d.IdConta) {
		errors["id_conta"] = "esse campo não pode ser vazio"
	}

	if d.TipoOperacao < 1 || d.TipoOperacao > 2 {
		errors["tipo_operacao"] = "tipo de operação inválido"
	}

	return errors
}
