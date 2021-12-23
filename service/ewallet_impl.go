package service

import (
	"github.com/pauluswi/bigben/model"
	"github.com/pauluswi/bigben/repository"
)

func NewEWalletService(accountRepository *repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		AccountRepository: *accountRepository,
	}
}

type ewalletServiceImpl struct {
	AccountRepository repository.AccountRepository
}

func (service *accountServiceImpl) EWalletTransfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error) {
	err = service.AccountRepository.Update(fromAccountNumber, toAccountNumber, amount)
	if err != nil {
		return err
	}
	return nil
}

func (service *accountServiceImpl) GetEWalletBalance(accountNumber int32) (response *model.AccountBalanceResponse, err error) {
	balance, err := service.AccountRepository.Find(accountNumber)
	if err != nil {
		return nil, err
	}

	response = &model.AccountBalanceResponse{
		AccountNumber: balance.AccountNumber,
		CustomerName:  balance.Name,
		Balance:       balance.Balance,
	}
	return response, nil
}
