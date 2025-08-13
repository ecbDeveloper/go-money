package services

import (
	"context"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewAccountService(pool *pgxpool.Pool) AccountService {
	return AccountService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, clientId uuid.UUID) (uuid.UUID, error) {
	accountId, err := a.queries.CreateAccount(ctx, clientId)
	if err != nil {
		return uuid.UUID{}, err
	}

	return accountId, nil
}
