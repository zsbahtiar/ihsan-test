package module

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zsbahtiar/ihsan-test/internal/core/dto"
	"github.com/zsbahtiar/ihsan-test/internal/core/entity"
	"github.com/zsbahtiar/ihsan-test/internal/core/repository"
	"math/rand"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

type AccountUsecase interface {
	RegisterCustomer(ctx context.Context, req dto.RegisterCustomerRequest) (dto.RegisterCustomerResponse, error)
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
