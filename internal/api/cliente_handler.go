package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/ecbDeveloper/go-money/internal/services"
)

func (api *Api) handleCreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.CreateClient
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, "requisição invalida", http.StatusBadRequest)
		return
	}

	validationErrs := client.Validate()
	if len(validationErrs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrs)
		return
	}

	clientId, err := api.UserService.CreateClient(r.Context(), client)
	if err != nil {
		slog.Error("failed to create client", "error", err)

		switch {
		case errors.Is(err, services.ErrDuplicateEmail):
			http.Error(w, "email já está em uso", http.StatusConflict)
			return

		case errors.Is(err, services.ErrDuplicateCPF):
			http.Error(w, "CPF já está em uso", http.StatusConflict)
			return

		case errors.Is(err, services.ErrDuplicateCNPJ):
			http.Error(w, "CNPJ já está em uso", http.StatusConflict)
			return

		case errors.Is(err, services.ErrInvalidCategory):
			http.Error(w, "categoria inválida", http.StatusBadRequest)
			return

		default:
			http.Error(w, "erro ao criar conta", http.StatusInternalServerError)
			return
		}
	}

	response := map[string]any{
		"mensagem":   "conta criada com sucesso",
		"id_cliente": clientId,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
