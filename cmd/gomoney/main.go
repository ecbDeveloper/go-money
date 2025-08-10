package main

import (
	"net/http"

	"github.com/ecbDeveloper/go-money/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	_ = csrf.ErrBadOrigin

	api := api.Api{
		Router: chi.NewRouter(),
	}

	api.BindRoutes()

	if err := http.ListenAndServe(":8082", api.Router); err != nil {
		panic(err)
	}
}
