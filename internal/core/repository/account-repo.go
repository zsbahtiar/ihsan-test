package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zsbahtiar/ihsan-test/internal/core/entity"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/database"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/response"
)

type accountRepository struct {
	db database.Postgres
}

type AccountRepository interface {
	CreateCustomer(ctx context.Context, customer entity.Customer, account entity.Account) error
}

func NewAccountRepository(db database.Postgres) AccountRepository {
	return &accountRepository{db: db}
}

const (
	queryCreateCustomer = `INSERT INTO customers(uuid, name, identity_number, phone_number) 
                          VALUES($1, $2, $3, $4) RETURNING id`
	queryCreateAccount = `INSERT INTO accounts(uuid, customer_id, account_number) 
                          VALUES($1, $2, $3)`
)

func (a *accountRepository) CreateCustomer(ctx context.Context, customer entity.Customer, account entity.Account) error {
	tx, err := a.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	err = tx.QueryRow(ctx, queryCreateCustomer, customer.Uuid, customer.Name, customer.IdentityNumber, customer.PhoneNumber).Scan(&customer.Id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.ConstraintName {
			case "customers_identity_number_key":
				return response.ErrDuplicateIdentityNumber
			case "customers_phone_number_key":
				return response.ErrDuplicatePhoneNumber
			}
		}
		return err
	}
	account.CustomerId = customer.Id

	_, err = tx.Exec(ctx, queryCreateAccount, account.Uuid, account.CustomerId, account.AccountNumber)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.ConstraintName == "accounts_account_number_key" {
				return response.ErrDuplicateAccountNumber
			}
		}
		return err
	}

	return tx.Commit(ctx)
}
