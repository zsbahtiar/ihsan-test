package dto

type RegisterCustomerResponse struct {
	NoRekening string `json:"no_rekening"`
}

type DepositResponse struct {
	Saldo float64 `json:"saldo"`
}

type WithdrawResponse struct {
	Saldo float64 `json:"saldo"`
}
