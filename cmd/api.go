package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pauluswi/bigben/config"
	"github.com/pauluswi/bigben/controller"
	"github.com/pauluswi/bigben/exception"
	"github.com/pauluswi/bigben/repository"
	"github.com/pauluswi/bigben/service"
	"github.com/spf13/cobra"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pauluswi/bigben/middleware"
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
			ewalletRepository := repository.NewEWalletRepository(database)

			// Setup Service
			ewalletService := service.NewEWalletService(&ewalletRepository)

			// Setup Controller
			ewalletController := controller.NewEWalletController(&ewalletService)

			// Setup Fiber
			app := fiber.New(config.NewFiberConfig())
			//app.Use(recover.New())

			app.Use(
				logger.New(), // add Logger middleware
			)

			app.Use(func(c *fiber.Ctx) error {
				// Set some security headers:
				c.Set("X-XSS-Protection", "1; mode=block")
				c.Set("X-Content-Type-Options", "nosniff")
				c.Set("X-Download-Options", "noopen")
				c.Set("Strict-Transport-Security", "max-age=5184000")
				c.Set("X-Frame-Options", "SAMEORIGIN")
				c.Set("X-DNS-Prefetch-Control", "off")

				// Go to next middleware:
				return c.Next()
			})

			// to validate request header
			secretKey := configuration.Get("secretKey")
			app.Use(middleware.New(middleware.Config{Key: secretKey}))

			// Setup Routing
			ewalletController.Route(app)

			// Start App
			err := app.Listen(":3000")
			exception.PanicIfNeeded(err)
		},
	}

	rootCmd.AddCommand(serveCmd)
}
