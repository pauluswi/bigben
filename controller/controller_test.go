package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pauluswi/bigben/config"
	"github.com/pauluswi/bigben/repository"
	"github.com/pauluswi/bigben/service"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	ewalletController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMySQLDatabase(configuration)
var accountRepository = repository.NewEWalletRepository(database)
var accountService = service.NewEWalletService(&accountRepository)

var ewalletController = NewEWalletController(&accountService)

var app = createTestApp()
