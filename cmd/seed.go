package cmd

import (
	"fmt"
	"os"

	"github.com/pauluswi/bigben/config"
	"github.com/pauluswi/bigben/database/seeder"
	"github.com/spf13/cobra"
)

var seedCmd *cobra.Command

func init() {
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "seeding dummy data",
		Long:  `Command to seeding dummy data`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running seed command")
			configuration := config.New()
			database := config.NewMySQLDatabase(configuration)
			seeder.Execute(database, args...)
			os.Exit(0)
		},
	}

	rootCmd.AddCommand(seedCmd)
}
