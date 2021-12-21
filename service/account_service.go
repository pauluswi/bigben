package service

import "github.com/pauluswi/bigben/model"

type AccountService interface {
	GetBalance(accountNumber int32) (response *model.AccountBalanceResponse, err error)
	Transfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error)
}
