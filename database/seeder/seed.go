package seeder

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Seed struct {
	Database *sql.DB
}

func Execute(db *sql.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(&s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		fmt.Println(seedType.NumMethod())
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(&s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}