package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/ecbDeveloper/go-money/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clientId, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	accountId, err := api.AccountService.CreateAccount(r.Context(), clientId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro interno inesperado no servidor",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"mensagem": "conta bancária criada com sucesso",
		"id_conta": accountId,
	})
}

func (api *Api) handleGetAccountBalanceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clientId, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	rawAccountId := chi.URLParam(r, "accountId")

	accountId, err := uuid.Parse(rawAccountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato inválido para o id da conta",
		})
		return
	}

	balance, err := api.AccountService.GetAccountBalanceById(r.Context(), accountId, clientId)
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFoundedOrNotOwned) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro interno inesperado no servidor",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"saldo": fmt.Sprintf("Saldo Atual: %.2f", balance),
	})
}

func (api *Api) handleAccountTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var accountTransactionRequest models.AccountTransaction
	err := json.NewDecoder(r.Body).Decode(&accountTransactionRequest)
	if err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, "requisição inválida", http.StatusBadRequest)
		return
	}

	accountId, err := uuid.Parse(accountTransactionRequest.IdConta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato do id da conta inválido",
		})
	}

	clientId, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		slog.Error("failed to get authenticated client id")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "acesso negado: faça login para continuar",
		})
		return
	}

	err = api.AccountService.AccountTransaction(r.Context(), accountId, clientId, accountTransactionRequest.Valor, accountTransactionRequest.TipoOperacao)
	if err != nil {
		slog.Error("failed to deposit money on client account", "error", err)

		if errors.Is(err, services.ErrAccountNotFoundedOrNotOwned) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		if errors.Is(err, services.ErrInvalidOperation) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro interno inesperado no servidor",
		})
		return
	}

	message := map[int32]string{
		1: "depósito realizado com sucesso",
		2: "saque realizado com sucesso",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": message[accountTransactionRequest.TipoOperacao],
	})
}
