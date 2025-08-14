package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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
			"error": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	accountId, err := api.AccountService.CreateAccount(r.Context(), clientId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro interno inesperado no servidor",
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
			"error": "erro inesperado, tente novamente mais tarde",
		})
		return
	}

	rawAccountId := chi.URLParam(r, "accountId")

	accountId, err := uuid.Parse(rawAccountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "formato inválido para o id da conta",
		})
		return
	}

	balance, err := api.AccountService.GetAccountBalanceById(r.Context(), accountId, clientId)
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFoundedOrNotOwned) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro interno inesperado no servidor",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"saldo": fmt.Sprintf("Saldo Atual: %.2f", balance),
	})
}
