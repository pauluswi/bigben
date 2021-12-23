package service

import (
	"github.com/pauluswi/bigben/model"
	"github.com/pauluswi/bigben/repository"
)

func NewEWalletService(ewalletRepository *repository.EWalletRepository) EWalletService {
	return &ewalletServiceImpl{
		EWalletRepository: *ewalletRepository,
	}
}

type ewalletServiceImpl struct {
	EWalletRepository repository.EWalletRepository
}

func (service *ewalletServiceImpl) EWalletTransfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error) {
	err = service.EWalletRepository.Update(fromAccountNumber, toAccountNumber, amount)
	if err != nil {
		return err
	}
	return nil
}

func (service *ewalletServiceImpl) GetEWalletBalance(accountNumber int32) (response *model.AccountBalanceResponse, err error) {
	balance, err := service.EWalletRepository.Find(accountNumber)
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
