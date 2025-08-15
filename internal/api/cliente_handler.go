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
	w.Header().Set("Content-Type", "application/json")

	var client models.CreateClient
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, "requisição invalida", http.StatusBadRequest)
		return
	}

	validationErrs := client.Validate()
	if len(validationErrs) > 0 {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrs)
		return
	}

	clientId, err := api.ClientService.CreateClient(r.Context(), client)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (api *Api) handleLoginClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials models.AuthenticateClient
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, "requisição inválida", http.StatusBadRequest)
		return
	}

	errs := credentials.Validate()
	if len(errs) > 0 {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	clientId, err := api.ClientService.AuthenticateClient(r.Context(), credentials.Email, credentials.Senha)
	if err != nil {
		slog.Error("failed to authenticate client", "error", err)
		if errors.Is(err, services.ErrInvalidCredentials) {
			http.Error(w, "credenciais incorretas, tente novamente", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro interno inesperado no servidor",
		})
		return
	}

	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		slog.Error("failed to renew session token", "error", err)
		http.Error(w, "erro interno inesperado no servidor", http.StatusInternalServerError)
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedClient", clientId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "cliente logado com sucesso",
	})
}
