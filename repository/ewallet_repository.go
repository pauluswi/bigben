package repository

import (
	"github.com/pauluswi/bigben/entity"
)

type EWalletRepository interface {
	Find(accountNumber int32) (balance *entity.Ewallet, err error)
	FindTransactions(accountNumber int32) (trxList []entity.EwalletTrx, err error)

	Transfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error)
	Deposit(toAccountNumber int32, amount int) (err error)
	Withdrawal(FromAccountNumber int32, amount int) (err error)
}
