package dto

type RegisterCustomerRequest struct {
	Nama string `json:"nama"`
	NoHp string `json:"no_hp"`
	Nik  string `json:"nik"`
}

type DepositRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"` // @Todo: for balance float is bad: change
}

type WithdrawRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"` // @Todo: for balance float is bad: change
}
