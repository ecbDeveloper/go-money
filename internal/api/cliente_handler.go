package api

import (
	"net/http"

	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/ecbDeveloper/go-money/internal/shared"
)

func (api *Api) handleCreateClient(w http.ResponseWriter, r *http.Request) {
	req, problems, err := shared.DecodeValidJson[models.CreateClient](r)
	if err != nil {
		shared.EncodeJson(w, http.StatusInternalServerError, problems)
		return
	}

	userId, err := api.UserService.CreateUsuarioAndConta(r.Context(), req)
	if err != nil {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to create user and bank account",
		})
		return
	}

	shared.EncodeJson(w, http.StatusCreated, map[string]any{
		"message": "account created successfully",
		"user_id": userId,
	})
}
