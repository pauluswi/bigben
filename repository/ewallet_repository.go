package repository

import (
	"github.com/pauluswi/bigben/entity"
)

type EWalletRepository interface {
	Find(accountNumber int32) (balance *entity.Ewallet, err error)
	Update(fromAccountNumber int32, toAccountNumber int32, amount int) (err error)
}
