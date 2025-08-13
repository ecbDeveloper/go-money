package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ecbDeveloper/go-money/internal/api"
	"github.com/ecbDeveloper/go-money/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *services.NewUserService(pool),
	}

	api.BindRoutes()

	log.Println("api started on port :8082")
	if err := http.ListenAndServe(":8082", api.Router); err != nil {
		panic(err)
	}
}
