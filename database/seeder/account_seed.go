package seeder

import "math/rand"

func (s *Seed) AccountSeed() {
	for i := 0; i < 10; i++ {
		stmt, _ := s.Database.Prepare(`INSERT INTO account (account_number, customer_number, balance) VALUES (?, ?, ?)`)

		min := 10000
		max := 30000

		_, err := stmt.Exec(55500+i, 1000+i, rand.Intn(max - min) + min)
		if err != nil {
			panic(err)
		}
	}
}
