package api

import (
	"encoding/json"
	"net/http"

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
