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

	clientID, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	accountID, err := api.AccountService.CreateAccount(r.Context(), clientID)
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
		"id_conta": accountID,
	})
}

func (api *Api) handleGetAccountBalanceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clientID, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	rawAccountID := chi.URLParam(r, "accountID")

	accountID, err := uuid.Parse(rawAccountID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato inválido para o id da conta",
		})
		return
	}

	balance, err := api.AccountService.GetAccountBalanceByID(r.Context(), accountID, clientID)
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

	validationErrs := accountTransactionRequest.Validate()
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrs)
		return
	}

	accountID, err := uuid.Parse(accountTransactionRequest.IDConta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato do id da conta inválido",
		})
	}

	clientID, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		slog.Error("failed to get authenticated client id")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	err = api.AccountService.AccountTransaction(r.Context(), accountID, clientID, accountTransactionRequest.Valor, accountTransactionRequest.TipoOperacao)
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

func (api *Api) handleMoneyTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transferMoneyRequest models.TransferMoney
	err := json.NewDecoder(r.Body).Decode(&transferMoneyRequest)
	if err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, "requisição inválida", http.StatusBadRequest)
		return
	}

	validationErrs := transferMoneyRequest.Validate()
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrs)
		return
	}

	destinyAccountID, err := uuid.Parse(transferMoneyRequest.IDContaDestino)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato do id da conta de destino inválido",
		})
		return
	}

	originAccountID, err := uuid.Parse(transferMoneyRequest.IDContaOrigem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato do id da conta de origem inválido",
		})
		return
	}

	clientID, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		slog.Error("failed to get authenticated client id")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	err = api.AccountService.MoneyTransfer(r.Context(), destinyAccountID, originAccountID, clientID, transferMoneyRequest.Valor)
	if err != nil {
		slog.Error("failed to make transfer", "error", err)

		if errors.Is(err, services.ErrAccountNotFounded) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": "não foi possível encontrar conta de destino",
			})
			return
		}

		if errors.Is(err, services.ErrCantTransferToSameAccount) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		if errors.Is(err, services.ErrAccountNotFoundedOrNotOwned) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		if errors.Is(err, services.ErrInsufficientBalance) {
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
		"mensagem": "transferência realizada com sucesso",
	})
}

func (api *Api) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rawAccountID := chi.URLParam(r, "accountID")

	accountID, err := uuid.Parse(rawAccountID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "formato do id da conta de destino inválido",
		})
		return
	}

	clientID, ok := api.Sessions.Get(r.Context(), "AuthenticatedClient").(uuid.UUID)
	if !ok {
		slog.Error("failed to get authenticated client id")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	err = api.AccountService.DeleteAccount(r.Context(), accountID, clientID)
	if err != nil {
		slog.Error("failed to delete account", "error", err)

		if errors.Is(err, services.ErrAccountNotFoundedOrNotOwned) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": err.Error(),
			})
			return
		}

		if errors.Is(err, services.ErrBalanceGreaterThenZero) {
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
		"mensagem": "conta deletada com sucesso",
	})
}
