package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/csrf"
)

func (api *Api) authMiddeware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "must be logged in",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *Api) handleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
