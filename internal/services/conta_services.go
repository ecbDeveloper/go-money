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

func (a *AccountService) CreateAccount(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	accountID, err := a.queries.CreateAccount(ctx, clientID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return accountID, nil
}

func (a *AccountService) GetAccountBalanceByID(ctx context.Context, accountID, clientID uuid.UUID) (float64, error) {
	clientAccounts, err := a.queries.GetAllAccountsByClientId(ctx, clientID)
	if err != nil {
		return 0, err
	}

	accountFounded := false
	for _, account := range clientAccounts {
		if account.IDCliente == clientID {
			accountFounded = true
			break
		}
	}

	if !accountFounded {
		return 0, ErrAccountNotFoundedOrNotOwned
	}

	balance, err := a.queries.GetBalanceByAccountId(ctx, accountID)
	if err != nil {
		return 0, err
	}

	balanceF, err := balance.Float64Value()
	if err != nil || !balanceF.Valid {
		return 0, err
	}

	return balanceF.Float64, nil
}

func (a *AccountService) AccountTransaction(ctx context.Context, accountID, clientID uuid.UUID, value float64, operationType int32) error {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := a.queries.WithTx(tx)

	account, err := queries.GetAccountForUpdate(ctx, accountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAccountNotFoundedOrNotOwned
		}
		return err
	}

	if account.IDCliente != clientID {
		return ErrAccountNotFoundedOrNotOwned
	}

	actualAccountBalanceF, err := shared.ConvertNumericToFloat(account.Saldo)
	if err != nil {
		return err
	}

	if operationType == 2 && (actualAccountBalanceF <= 0 || value > actualAccountBalanceF) {
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
		ID:    accountID,
	}

	transferArgs := sqlc.CreateTransferenciaParams{
		IDConta: accountID,
		Valor:   valueNumeric,
		Tipo:    operationType,
	}

	err = queries.UpdateAccountBalance(ctx, depositArgs)
	if err != nil {
		return err
	}

	err = queries.CreateTransferencia(ctx, transferArgs)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (a *AccountService) MoneyTransfer(ctx context.Context, destinyAccountID, originAccountID, clientID uuid.UUID, value float64) error {
	if originAccountID == destinyAccountID {
		return ErrCantTransferToSameAccount
	}

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := a.queries.WithTx(tx)

	originAccount, err := queries.GetAccountForUpdate(ctx, originAccountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAccountNotFoundedOrNotOwned
		}
		return err
	}

	if originAccount.IDCliente != clientID {
		return ErrAccountNotFoundedOrNotOwned
	}

	destinyAccount, err := queries.GetAccountForUpdate(ctx, destinyAccountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAccountNotFounded
		}
		return err
	}

	destinyActualBalanceF, err := shared.ConvertNumericToFloat(destinyAccount.Saldo)
	if err != nil {
		return err
	}

	originActualBalanceF, err := shared.ConvertNumericToFloat(originAccount.Saldo)
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

	err = queries.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:    destinyAccountID,
		Saldo: destinyNewBalance,
	})
	if err != nil {
		return err
	}

	err = queries.UpdateAccountBalance(ctx, sqlc.UpdateAccountBalanceParams{
		ID:    originAccountID,
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

func (a *AccountService) DeleteAccount(ctx context.Context, accountID, clientID uuid.UUID) error {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := a.queries.WithTx(tx)

	account, err := queries.GetAccountForUpdate(ctx, accountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAccountNotFoundedOrNotOwned
		}
		return err
	}

	if account.IDCliente != clientID {
		return ErrAccountNotFoundedOrNotOwned
	}

	actualBalanceF, err := shared.ConvertNumericToFloat(account.Saldo)
	if err != nil {
		return err
	}

	if actualBalanceF > 0 {
		return ErrBalanceGreaterThenZero
	}

	accountStatus := sqlc.UpdateAccountStatusParams{
		Status: pgtype.Int4{
			Int32: 2,
			Valid: true,
		},
		ID: accountID,
	}

	if err := queries.UpdateAccountStatus(ctx, accountStatus); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
