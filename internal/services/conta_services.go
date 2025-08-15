package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/ecbDeveloper/go-money/internal/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrAccountNotFoundedOrNotOwned = errors.New("conta não encontrada ou não pertence ao cliente")
	ErrAccountNotFounded           = errors.New("conta de destino não encontrada")
	ErrInvalidOperation            = errors.New("operação inválida")
	ErrCantTransferToSameAccount   = errors.New("não é possível transferir dinheiro pra própria conta")
	ErrInsufficientBalance         = errors.New("saldo insuficiente para realizar a transação")
	ErrBalanceGreaterThenZero      = errors.New("há saldo na sua conta, saque ou transfira-o para deletar a conta")
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

	actualAccountBalanceF, err := shared.ConvertNumericToFloat(actualAccountBalance)
	if err != nil {
		return err
	}

	if operationType == 1 && (actualAccountBalanceF <= 0 || value > actualAccountBalanceF) {
		return ErrInsufficientBalance
	}

	valueNumeric, err := shared.ConvertFloatToNumeric(value)
	if err != nil {
		return err
	}

	var newBalanceNumeric pgtype.Numeric
	switch operationType {
	case 1:
		newBalanceNumeric, err = shared.ConvertFloatToNumeric(actualAccountBalanceF + value)
		if err != nil {
			return err
		}
	case 2:
		newBalanceNumeric, err = shared.ConvertFloatToNumeric(actualAccountBalanceF - value)
		if err != nil {
			return err
		}
	default:
		return ErrInvalidOperation
	}

	depositArgs := sqlc.UpdateAccountBalanceParams{
		Saldo: newBalanceNumeric,
		ID:    accountId,
	}

	transferArgs := sqlc.CreateTransferenciaParams{
		IDConta: accountId,
		Valor:   valueNumeric,
		Tipo:    operationType,
	}

	err = queries.UpdateAccountBalance(ctx, depositArgs)
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

func (a *AccountService) MoneyTransfer(ctx context.Context, destinyAccountId, originAccountId, clientId uuid.UUID, value float64) error {
	if originAccountId == destinyAccountId {
		return ErrCantTransferToSameAccount
	}

	clientAccounts, err := a.queries.GetAllAccountsByClientId(ctx, clientId)
	if err != nil {
		return err
	}

	accountFounded := false
	var originActualBalance pgtype.Numeric
	for _, account := range clientAccounts {
		if originAccountId == account.ID {
			originActualBalance = account.Saldo
			accountFounded = true
			break
		}
	}

	if !accountFounded {
		return ErrAccountNotFoundedOrNotOwned
	}

	destinyActualBalance, err := a.queries.GetBalanceByAccountId(ctx, destinyAccountId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAccountNotFounded
		}
		return err
	}

	destinyActualBalanceF, err := shared.ConvertNumericToFloat(destinyActualBalance)
	if err != nil {
		return err
	}

	originActualBalanceF, err := shared.ConvertNumericToFloat(originActualBalance)
	if err != nil {
		return err
	}

	if originActualBalanceF <= 0 || value > originActualBalanceF {
		return ErrInsufficientBalance
	}

	destinyNewBalance, err := shared.ConvertFloatToNumeric(destinyActualBalanceF + value)
	if err != nil {
		return err
	}

	originNewBalance, err := shared.ConvertFloatToNumeric(originActualBalanceF - value)
	if err != nil {
		return err
	}

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}

	queries := sqlc.New(tx)

	err = queries.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:    destinyAccountId,
		Saldo: destinyNewBalance,
	})
	if err != nil {
		return err
	}

	err = queries.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:    originAccountId,
		Saldo: originNewBalance,
	})
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (a *AccountService) DeleteAccount(ctx context.Context, accountId, clientId uuid.UUID) error {
	clientAccounts, err := a.queries.GetAllAccountsByClientId(ctx, clientId)
	if err != nil {
		return err
	}

	accountFounded := false
	var actualBalance pgtype.Numeric
	for _, account := range clientAccounts {
		if accountId == account.ID {
			actualBalance = account.Saldo
			accountFounded = true
			break
		}
	}

	if !accountFounded {
		return ErrAccountNotFoundedOrNotOwned
	}

	actualBalanceF, err := shared.ConvertNumericToFloat(actualBalance)
	if err != nil {
		return err
	}

	if actualBalanceF > 0 {
		return ErrBalanceGreaterThenZero
	}

	accountStatus := sqlc.UpdateAccountStatusParams{
		Status: pgtype.Int4{
			Int32: 1,
			Valid: true,
		},
		ID: accountId,
	}

	if err := a.queries.UpdateAccountStatus(ctx, accountStatus); err != nil {
		return err
	}

	return nil
}
