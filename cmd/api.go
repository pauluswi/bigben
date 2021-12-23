package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pauluswi/bigben/config"
	"github.com/pauluswi/bigben/controller"
	"github.com/pauluswi/bigben/exception"
	"github.com/pauluswi/bigben/repository"
	"github.com/pauluswi/bigben/service"
	"github.com/spf13/cobra"
)

var serveCmd *cobra.Command

func init() {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve api",
		Long:  `Command to serve application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running serve command")

			// Setup Configuration
			configuration := config.New()
			database := config.NewMySQLDatabase(configuration)

			// Setup Repository
			accountRepository := repository.NewAccountRepository(database)
			ewalletRepository := repository.NewEWalletRepository(database)

			// Setup Service
			accountService := service.NewAccountService(&accountRepository)
			ewalletService := service.NewEWalletService(&ewalletRepository)

			// Setup Controller
			accountController := controller.NewAccountController(&accountService)
			ewalletController := controller.NewEWalletController(&ewalletService)

			// Setup Fiber
			app := fiber.New(config.NewFiberConfig())
			app.Use(recover.New())

			// Setup Routing
			accountController.Route(app)
			ewalletController.Route(app)

			// Start App
			err := app.Listen(":3000")
			exception.PanicIfNeeded(err)
		},
	}

	rootCmd.AddCommand(serveCmd)
}
