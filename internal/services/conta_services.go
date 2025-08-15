package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/ecbDeveloper/go-money/internal/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrAccountNotFoundedOrNotOwned = errors.New("conta não encontrada ou não pertence ao cliente")
	ErrInvalidOperation            = errors.New("operação inválida")
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

func (a *AccountService) AccountTransaction(ctx context.Context, accountId, clientId uuid.UUID, value float64, operationType int32) error {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}

	queries := sqlc.New(tx)

	clientAccounts, err := queries.GetAllAccountsByClientId(ctx, clientId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	accountFounded := false
	var actualAccountBalance pgtype.Numeric
	for _, account := range clientAccounts {
		if accountId == account.ID {
			actualAccountBalance = account.Saldo
			accountFounded = true
			break
		}
	}

	if !accountFounded {
		return ErrAccountNotFoundedOrNotOwned
	}

	actualAccountBalanceFloat, err := shared.ConvertNumericToFloat(actualAccountBalance)
	if err != nil {
		return err
	}

	valueNumeric, err := shared.ConvertFloatToNumeric(value)
	if err != nil {
		return err
	}

	var newBalanceNumeric pgtype.Numeric
	switch operationType {
	case 1:
		newBalanceNumeric, err = shared.ConvertFloatToNumeric(actualAccountBalanceFloat + value)
		if err != nil {
			return err
		}
	case 2:
		newBalanceNumeric, err = shared.ConvertFloatToNumeric(actualAccountBalanceFloat - value)
		if err != nil {
			return err
		}
	default:
		return ErrInvalidOperation
	}

	depositArgs := sqlc.PutMoneyInAccountParams{
		Saldo: newBalanceNumeric,
		ID:    accountId,
	}

	transferArgs := sqlc.CreateTransferenciaParams{
		IDConta: accountId,
		Valor:   valueNumeric,
		Tipo:    operationType,
	}

	err = queries.PutMoneyInAccount(ctx, depositArgs)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = queries.CreateTransferencia(ctx, transferArgs)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
