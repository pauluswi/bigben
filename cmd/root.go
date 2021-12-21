package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd  = &cobra.Command{
	Use : "cli",
	Short: "Root command for our application",
	Long:  `Root command for our application, the main purpose is to help setup subcommands`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}