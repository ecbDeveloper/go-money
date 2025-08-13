package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCategory = errors.New("you need to select a valid category")
	ErrDuplicateEmail  = errors.New("this email is already in use")
	ErrDuplicateCPF    = errors.New("this cpf is already in use")
	ErrDuplicateCNPJ   = errors.New("this cnpj is already in use")
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (u *UserService) CreateClient(ctx context.Context, client models.CreateClient) (uuid.UUID, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	queries := sqlc.New(tx)

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

	clientId, err := queries.CreateClient(ctx, arguments)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return uuid.UUID{}, ErrDuplicateEmail
			}
		}
		tx.Rollback(ctx)
		return uuid.UUID{}, err
	}

	switch client.Categoria {
	case 1:
		arguments := sqlc.CreatePessoaFisicaParams{
			IDCliente:      clientId,
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

			tx.Rollback(ctx)
			return uuid.UUID{}, err
		}

	case 2:
		arguments := sqlc.CreatePessoaJuridicaParams{
			IDCliente:    clientId,
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

			tx.Rollback(ctx)
			return uuid.UUID{}, err
		}

	default:
		return uuid.UUID{}, ErrInvalidCategory
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, err
	}

	return clientId, nil
}
