package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *API) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("CSRF_SECRET")),
		csrf.Secure(false),
	)

	api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrf", api.handleGetCSRFToken)

			r.Post("/client", api.handleCreateClient)
			r.Post("/client/login", api.handleLoginClient)

			r.Group(func(r chi.Router) {
				r.Use(api.authMiddeware)
				r.Get("/account/{accountId}/balance", api.handleGetAccountBalanceByID)
				r.Post("/account", api.handleCreateAccount)
				r.Post("/account/transfer", api.handleMoneyTransfer)
				r.Post("/account/transaction", api.handleAccountTransaction)
				r.Delete("/account/{accountId}", api.handleDeleteAccount)
			})
		})
	})
}
