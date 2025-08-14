package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrAccountNotFoundedOrNotOwned = errors.New("conta não encontrada ou não pertence ao cliente")
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

func (a *AccountService) GetAccountBalanceById(ctx context.Context, accountId, clientId uuid.UUID) (float64, error) {
	clientAccounts, err := a.queries.GetAllAccountsByClientId(ctx, clientId)
	if err != nil {
		return 0, err
	}

	accountFounded := false
	for _, account := range clientAccounts {
		if account.IDCliente == clientId {
			accountFounded = true
			break
		}
	}

	if !accountFounded {
		return 0, ErrAccountNotFoundedOrNotOwned
	}

	balance, err := a.queries.GetBalanceByAccountId(ctx, accountId)
	if err != nil {
		return 0, err
	}

	balanceF, err := balance.Float64Value()
	if err != nil || !balanceF.Valid {
		return 0, err
	}

	return balanceF.Float64, nil
}
