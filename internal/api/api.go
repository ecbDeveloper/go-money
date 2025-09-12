package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/ecbDeveloper/go-money/internal/services"
	"github.com/go-chi/chi/v5"
)

type API struct {
	Router         *chi.Mux
	Sessions       *scs.SessionManager
	ClientService  services.ClientService
	AccountService services.AccountService
}
