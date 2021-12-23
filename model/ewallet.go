package model

type CreateEWalletTransferRequest struct {
	ToAccountNumber int32 `json:"to_account_number"`
	Amount          int   `json:"amount"`
}

type EWalletBalanceResponse struct {
	AccountNumber int32  `json:"account_number"`
	Balance       int    `json:"balance"`
	ModifiedDate  string `json:"modified_date"`
}
