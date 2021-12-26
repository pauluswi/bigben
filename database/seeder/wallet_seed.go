package seeder

func (s *Seed) WalletSeed() {
	for i := 0; i < 5; i++ {
		stmt, _ := s.Database.Prepare(`INSERT INTO db_wallet.wallet(accountid, balance, modifieddate) VALUES (?, ?, Now()`)

		_, err := stmt.Exec(10000+i, 100000)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Seed) WalletTrxSeed() {
	for i := 0; i < 5; i++ {
		stmt, _ := s.Database.Prepare(`INSERT INTO db_wallet.wallet_trx(accountid, trxtype, dc, trxamount, createdby, createddate, modifiedby, modifieddate) VALUES (?, ?, ?, ?, ?, Now(), ?, Now()`)

		_, err := stmt.Exec(10000+i, "INIT", "C", 100000, "user", "user")
		if err != nil {
			panic(err)
		}
	}
}
