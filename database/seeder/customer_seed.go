package seeder

import (
	"github.com/bxcodec/faker/v3"
)

func (s *Seed) CustomerSeed() {
	for i := 0; i < 10; i++ {
		stmt, _ := s.Database.Prepare(`INSERT INTO customer (customer_number, name) VALUES (?, ?)`)
		_, err := stmt.Exec(1000+i, faker.Name())
		if err != nil {
			panic(err)
		}
	}
}
