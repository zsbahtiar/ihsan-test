package entity

type Customer struct {
	Id             int64
	Uuid           string
	Name           string
	IdentityNumber string
	PhoneNumber    string
}

type Account struct {
	Id            int64   `db:"id"`
	Uuid          string  `db:"uuid"`
	CustomerId    int64   `db:"customer_id"`
	AccountNumber string  `db:"account_number"`
	Balance       float64 `db:"balance"` //@Todo: for currency float is bad, change
}

type TransactionType string

const (
	TransactionDeposit  TransactionType = "DEPOSIT"
	TransactionWithdraw                 = "WITHDRAW"
)

type Transaction struct {
	Id              int64
	Uuid            string
	AccountId       int64
	TransactionType TransactionType
	Amount          float64
}
