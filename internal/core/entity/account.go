package entity

type Customer struct {
	Id             int64
	Uuid           string
	Name           string
	IdentityNumber string
	PhoneNumber    string
}

type Account struct {
	Id            int64
	Uuid          string
	CustomerId    int64
	AccountNumber string
	Balance       float64 //@Todo: for currency float is bad, change
}
