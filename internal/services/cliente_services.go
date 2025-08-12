package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-money/internal/db/sqlc"
	"github.com/ecbDeveloper/go-money/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidCategory = errors.New("you need to select a valid category")
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

func (u *UserService) CreateUsuarioAndConta(ctx context.Context, client models.CreateClient) (uuid.UUID, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	queries := sqlc.New(tx)

	arguments := sqlc.CreateClientParams{
		CategoriaCliente: client.Categoria,
		Telefone:         client.Telefone,
		Email:            client.Email,
	}

	clientId, err := queries.CreateClient(ctx, arguments)
	if err != nil {
		// TODO VALIDAR SE TELEFONE E EMAIL E UNIQUE
		tx.Rollback(ctx)
		return uuid.UUID{}, err
	}

	switch client.Categoria.Int32 {
	case 1:
		arguments := sqlc.CreatePessoaFisicaParams{
			IDCliente:      clientId,
			NomeCompleto:   client.PessoaFisica.NomeCompleto,
			DataNascimento: client.PessoaFisica.DataNascimento,
			Cpf:            client.PessoaFisica.Cpf,
		}

		err = queries.CreatePessoaFisica(ctx, arguments)
		if err != nil {
			// TODO VALIDAR CPF SE É UNIQUE
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
			tx.Rollback(ctx)
			//TODO VALIDAR SE CNPJ E UNIQUE
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
