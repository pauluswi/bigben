package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/pauluswi/bigben/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1",
		Long:  `Command to downgrade database from v2 to v1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down command")
			configuration := config.New()
			database := config.NewMySQLDatabase(configuration)
			dbDriver, err := mysql.WithInstance(database, &mysql.Config{})
			if err != nil {
				fmt.Printf("instance error: %v \n", err)
			}

			fileSource, err := (&file.File{}).Open("file://database/migrations")
			if err != nil {
				fmt.Printf("opening file error: %v \n", err)
			}

			m, err := migrate.NewWithInstance("file", fileSource, configuration.Get("MYSQL_DB_NAME"), dbDriver)
			if err != nil {
				fmt.Printf("migrate error: %v \n", err)
			}

			if err = m.Down(); err != nil {
				fmt.Printf("migrate down error: %v \n", err)
			}

			fmt.Println("Migrate down done with success")
		},
	}

	cmd.AddCommand(migrateDownCmd)
}
