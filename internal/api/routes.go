package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, api.Sessions.LoadAndSave)
	_ = csrf.Protect(
		[]byte(os.Getenv("CSRF_SECRET")),
		csrf.Secure(false),
	)

	api.Router.Route("/api/v1", func(r chi.Router) {
		r.Post("/client", api.handleCreateClient)
		r.Post("/client/login", api.handleLoginClient)
		r.Post("/account", api.handleCreateAccount)
		r.Get("/account/{accountId}/balance", api.handleGetAccountBalanceById)
		r.Post("/account/transaction", api.handleAccountTransaction)

		api.Router.Group(func(r chi.Router) {
		})

	})
}
