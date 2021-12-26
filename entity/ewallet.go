package entity

type Ewallet struct {
	AccountID    int32
	Balance      int
	ModifiedDate string
}

type EwalletTrx struct {
	TrxID        int32
	AccountID    int32
	TrxType      string
	CD           string
	Amount       int32
	CreatedBy    string
	CreatedDate  string
	ModifiedBy   string
	ModifiedDate string
}
