package module

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zsbahtiar/ihsan-test/internal/core/dto"
	"github.com/zsbahtiar/ihsan-test/internal/core/entity"
	"github.com/zsbahtiar/ihsan-test/internal/core/repository"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/response"
	"math/rand"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

type AccountUsecase interface {
	RegisterCustomer(ctx context.Context, req dto.RegisterCustomerRequest) (dto.RegisterCustomerResponse, error)
	Deposit(ctx context.Context, req dto.DepositRequest) (dto.DepositResponse, error)
	Withdraw(ctx context.Context, req dto.WithdrawRequest) (dto.WithdrawResponse, error)
	GetAccountDetail(ctx context.Context, accountNumber string) (dto.GetAccountDetailResponse, error)
}

func NewAccountUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{accountRepo: accountRepo}
}

func (a *accountUsecase) RegisterCustomer(ctx context.Context, req dto.RegisterCustomerRequest) (dto.RegisterCustomerResponse, error) {
	customer := entity.Customer{
		Uuid:           uuid.New().String(),
		Name:           req.Nama,
		IdentityNumber: req.Nik,
		PhoneNumber:    req.NoHp,
	}
	account := entity.Account{
		Uuid:          uuid.New().String(),
		AccountNumber: fmt.Sprintf("1234%06d", rand.Intn(1000000)),
	}

	err := a.accountRepo.CreateCustomer(ctx, customer, account)
	if err != nil {
		return dto.RegisterCustomerResponse{}, err
	}

	return dto.RegisterCustomerResponse{
		NoRekening: account.AccountNumber,
	}, err
}

func (a *accountUsecase) Deposit(ctx context.Context, req dto.DepositRequest) (dto.DepositResponse, error) {
	account, err := a.accountRepo.GetAccountByAccountNumber(ctx, req.NoRekening)
	if err != nil {
		return dto.DepositResponse{}, err
	}

	account.Balance += req.Nominal
	transaction := entity.Transaction{
		Uuid:            uuid.NewString(),
		AccountId:       account.Id,
		TransactionType: entity.TransactionDeposit,
		Amount:          req.Nominal,
	}

	err = a.accountRepo.CreateTransaction(ctx, transaction, account)
	if err != nil {
		return dto.DepositResponse{}, err
	}

	return dto.DepositResponse{Saldo: account.Balance}, nil

}

func (a *accountUsecase) Withdraw(ctx context.Context, req dto.WithdrawRequest) (dto.WithdrawResponse, error) {
	account, err := a.accountRepo.GetAccountByAccountNumber(ctx, req.NoRekening)
	if err != nil {
		return dto.WithdrawResponse{}, err
	}

	if req.Nominal > account.Balance {
		return dto.WithdrawResponse{}, response.ErrInsufficientBalance
	}

	account.Balance -= req.Nominal
	transaction := entity.Transaction{
		Uuid:            uuid.NewString(),
		AccountId:       account.Id,
		TransactionType: entity.TransactionWithdraw,
		Amount:          req.Nominal,
	}

	err = a.accountRepo.CreateTransaction(ctx, transaction, account)
	if err != nil {
		return dto.WithdrawResponse{}, err
	}

	return dto.WithdrawResponse{
		Saldo: account.Balance,
	}, nil

}

func (a *accountUsecase) GetAccountDetail(ctx context.Context, accountNumber string) (dto.GetAccountDetailResponse, error) {
	account, err := a.accountRepo.GetAccountByAccountNumber(ctx, accountNumber)
	if err != nil {
		return dto.GetAccountDetailResponse{}, err
	}

	return dto.GetAccountDetailResponse{
		NoRekening: account.AccountNumber,
		Saldo:      account.Balance,
	}, err
}
