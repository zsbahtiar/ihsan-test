package dto

type RegisterCustomerRequest struct {
	Nama string `json:"nama"`
	NoHp string `json:"no_hp"`
	Nik  string `json:"nik"`
}
