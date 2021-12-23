package entity

type Ewallet struct {
	AccountID    int32
	Balance      int
	ModifiedDate string
}

type EwalletTrx struct {
	AccountID    int32
	TrxID        int32
	TrxType      string
	CD           string
	Amount       int
	CreatedBy    string
	CreatedDate  string
	ModifiedBy   string
	ModifiedDate string
}
