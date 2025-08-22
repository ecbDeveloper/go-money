package models

import "github.com/ecbDeveloper/go-money/internal/shared"

type AccountTransaction struct {
	Valor        float64 `json:"valor"`
	IDConta      string  `json:"id_conta"`
	TipoOperacao int32   `json:"tipo_operacao"`
}

func (d AccountTransaction) Validate() map[string]string {
	errors := make(map[string]string)

	if d.Valor < 0.01 {
		errors["valor"] = "depósito precisa ser maior que R$0,00"
	}

	if shared.IsBlank(d.IDConta) {
		errors["id_conta"] = "esse campo não pode ser vazio"
	}

	if d.TipoOperacao < 1 || d.TipoOperacao > 2 {
		errors["tipo_operacao"] = "tipo de operação inválido"
	}

	return errors
}

type TransferMoney struct {
	Valor          float64 `json:"valor"`
	IDContaDestino string  `json:"id_conta_destino"`
	IDContaOrigem  string  `json:"id_conta_origem"`
}

func (d TransferMoney) Validate() map[string]string {
	errors := make(map[string]string)

	if d.Valor < 0.01 {
		errors["valor"] = "depósito precisa ser maior que R$0,00"
	}

	if shared.IsBlank(d.IDContaDestino) {
		errors["id_conta_destino"] = "esse campo não pode ser vazio"
	}

	if shared.IsBlank(d.IDContaOrigem) {
		errors["id_conta_origem"] = "esse campo não pode ser vazio"
	}

	return errors
}
