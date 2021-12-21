package cmd

import (
	"fmt"

	"github.com/pauluswi/bigben/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
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

			if err = m.Up(); err != nil {
				fmt.Printf("migrate up error: %v \n", err)
			}

			fmt.Println("Migrate up done with success")

		},
	}

	cmd.AddCommand(migrateUpCmd)
}
