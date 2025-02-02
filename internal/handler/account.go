package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zsbahtiar/ihsan-test/internal/core/dto"
	"github.com/zsbahtiar/ihsan-test/internal/core/module"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/response"
	"net/http"
)

type accountHandler struct {
	accountUsecase module.AccountUsecase
}

type AccountHandler interface {
	Register(ctx *fiber.Ctx) error
	Deposit(ctx *fiber.Ctx) error
	Withdraw(ctx *fiber.Ctx) error
	GetAccountDetail(ctx *fiber.Ctx) error
}

func NewAccountHandler(accountUsecase module.AccountUsecase) AccountHandler {
	return &accountHandler{accountUsecase: accountUsecase}
}

func (a *accountHandler) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Invalid request payload",
		})
	}
	data, err := a.accountUsecase.RegisterCustomer(ctx.Context(), req)
	if err != nil {
		return response.WriteError(ctx, err)
	}

	return response.WriteSuccess(ctx, http.StatusCreated, data)
}

func (a *accountHandler) Deposit(ctx *fiber.Ctx) error {
	var req dto.DepositRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Invalid request payload",
		})
	}
	data, err := a.accountUsecase.Deposit(ctx.Context(), req)
	if err != nil {
		return response.WriteError(ctx, err)
	}

	return response.WriteSuccess(ctx, http.StatusOK, data)
}

func (a *accountHandler) Withdraw(ctx *fiber.Ctx) error {
	var req dto.WithdrawRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Invalid request payload",
		})
	}
	data, err := a.accountUsecase.Withdraw(ctx.Context(), req)
	if err != nil {
		return response.WriteError(ctx, err)
	}

	return response.WriteSuccess(ctx, http.StatusOK, data)
}

func (a *accountHandler) GetAccountDetail(ctx *fiber.Ctx) error {
	accountNumber := ctx.Params("no_rekening")

	data, err := a.accountUsecase.GetAccountDetail(ctx.Context(), accountNumber)
	if err != nil {
		return response.WriteError(ctx, err)
	}

	return response.WriteSuccess(ctx, http.StatusOK, data)
}
