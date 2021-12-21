package repository

import (
	"github.com/pauluswi/bigben/entity"
)

type AccountRepository interface {
	Find(accountNumber int32) (balance *entity.Balance, err error)
	Update(fromAccountNumber int32, toAccountNumber int32, amount int) (err error)
}
