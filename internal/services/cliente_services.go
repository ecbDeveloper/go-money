package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCategory    = errors.New("you need to select a valid category")
	ErrDuplicateEmail     = errors.New("this email is already in use")
	ErrDuplicateCPF       = errors.New("this cpf is already in use")
	ErrDuplicateCNPJ      = errors.New("this cnpj is already in use")
	ErrInvalidCredentials = errors.New("email or password is incorrect")
)

type ClientService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewClientService(pool *pgxpool.Pool) ClientService {
	return ClientService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (c *ClientService) CreateClient(ctx context.Context, client models.CreateClient) (uuid.UUID, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer tx.Rollback(ctx)

	queries := c.queries.WithTx(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(client.Senha), 10)
	if err != nil {
		return uuid.UUID{}, err
	}

	arguments := sqlc.CreateClientParams{
		CategoriaCliente: pgtype.Int4{
			Int32: client.Categoria,
			Valid: true,
		},
		Telefone: client.Telefone,
		Email:    client.Email,
		Password: hashedPassword,
	}

	clientID, err := queries.CreateClient(ctx, arguments)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return uuid.UUID{}, ErrDuplicateEmail
			}
		}
		return uuid.UUID{}, err
	}

	switch client.Categoria {
	case 1:
		arguments := sqlc.CreatePessoaFisicaParams{
			IDCliente:      clientID,
			NomeCompleto:   client.PessoaFisica.NomeCompleto,
			DataNascimento: client.PessoaFisica.DataNascimento,
			Cpf:            client.PessoaFisica.Cpf,
		}

		err = queries.CreatePessoaFisica(ctx, arguments)
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgErr.Code == "23505" {
					return uuid.UUID{}, ErrDuplicateCPF
				}
			}

			return uuid.UUID{}, err
		}

	case 2:
		arguments := sqlc.CreatePessoaJuridicaParams{
			IDCliente:    clientID,
			DataCriacao:  client.PessoaJuridica.DataCriacao,
			NomeFantasia: client.PessoaJuridica.NomeFantasia,
			Cnpj:         client.PessoaJuridica.Cnpj,
		}

		err = queries.CreatePessoaJuridica(ctx, arguments)
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgErr.Code == "23505" {
					return uuid.UUID{}, ErrDuplicateCNPJ
				}
			}

			return uuid.UUID{}, err
		}

	default:
		return uuid.UUID{}, ErrInvalidCategory
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, err
	}

	return clientID, nil
}

func (c *ClientService) AuthenticateClient(ctx context.Context, email, password string) (uuid.UUID, error) {
	client, err := c.queries.GetClientByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrInvalidCredentials
		}

		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(client.Password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, ErrInvalidCredentials
		}

		return uuid.UUID{}, err
	}

	return client.ID, nil
}
