package entity

type Account struct {
	AccountNumber  int32
	CustomerNumber int32
	Balance        int
}

type Balance struct {
	AccountNumber int32
	Name          string
	Balance       int
}