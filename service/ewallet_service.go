package service

import "github.com/pauluswi/bigben/model"

type EWalletService interface {
	GetEWalletBalance(accountNumber int32) (response *model.EWalletBalanceResponse, err error)
	EWalletTransfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error)
	GetEWalletTransactions(accountNumber int32) (response []model.EWalletTransactionsResponse, err error)
}
