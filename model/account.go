package model

type CreateTransferRequest struct {
	ToAccountNumber int32 `json:"to_account_number"`
	Amount          int   `json:"amount"`
}

type AccountBalanceResponse struct {
	AccountNumber int32  `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}