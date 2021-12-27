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

type EWalletTransactionsResponse struct {
	AccountID    int32  `json:"account_number"`
	TrxID        int32  `json:"trx_id"`
	TrxType      string `json:"trx_type"`
	CD           string `json:"cd"`
	Amount       int32  `json:"amount"`
	CreatedBy    string `json:"created_by"`
	CreatedDate  string `json:"created_date"`
	ModifiedBy   string `json:"modified_by"`
	ModifiedDate string `json:"modified_date"`
}

type EWalletDepositRequest struct {
	ToAccountNumber int32 `json:"to_account_number"`
	Amount          int   `json:"amount"`
}

type EWalletWitdrawalRequest struct {
	FromAccountNumber int32 `json:"from_account_number"`
	Amount            int   `json:"amount"`
}
