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
	err = service.EWalletRepository.Transfer(fromAccountNumber, toAccountNumber, amount)
	if err != nil {
		return err
	}
	return nil
}

func (service *ewalletServiceImpl) GetEWalletBalance(accountNumber int32) (response *model.EWalletBalanceResponse, err error) {
	balance, err := service.EWalletRepository.Find(accountNumber)
	if err != nil {
		return nil, err
	}

	response = &model.EWalletBalanceResponse{
		AccountNumber: balance.AccountID,
		Balance:       balance.Balance,
		ModifiedDate:  balance.ModifiedDate,
	}
	return response, nil
}

func (service *ewalletServiceImpl) GetEWalletTransactions(accountNumber int32) (response []model.EWalletTransactionsResponse, err error) {
	rows, err := service.EWalletRepository.FindTransactions(accountNumber)
	if err != nil {
		return nil, err
	}

	var trxs []model.EWalletTransactionsResponse

	for i := 0; i < len(rows); i++ {
		var trx model.EWalletTransactionsResponse

		trx.AccountID = rows[i].AccountID
		trx.TrxID = rows[i].TrxID
		trx.TrxType = rows[i].TrxType
		trx.CD = rows[i].CD
		trx.Amount = rows[i].Amount
		trx.CreatedBy = rows[i].CreatedBy
		trx.CreatedDate = rows[i].CreatedDate
		trx.ModifiedBy = rows[i].ModifiedBy
		trx.ModifiedDate = rows[i].ModifiedDate

		trxs = append(trxs, trx)
	}

	return trxs, nil
}
