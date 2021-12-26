package controller

// GET http://localhost:3000/v1/ewallet/balance/{account_id}
// POST http://localhost:3000/v1/ewallet/deposit
// GET http://localhost:3000/v1/ewallet/transactions
// POST http://localhost:3000/v1/ewallet/transfer
// POST http://localhost:3000/v1/ewallet/withdrawal

import (
	"strconv"

	"github.com/pauluswi/bigben/exception"
	"github.com/pauluswi/bigben/model"
	"github.com/pauluswi/bigben/service"
	"github.com/pauluswi/bigben/validation"

	"github.com/gofiber/fiber/v2"
)

type EWalletController struct {
	EWalletService service.EWalletService
}

func NewEWalletController(ewalletService *service.EWalletService) EWalletController {
	return EWalletController{
		EWalletService: *ewalletService,
	}
}

func (controller *EWalletController) Route(app *fiber.App) {
	app.Get("/v1/ewallet/balance/:account_id", controller.GetBalance)
	app.Get("/v1/ewallet/transaction/history/:account_id", controller.GetTransactions)
	app.Post("/v1/ewallet/transaction/transfer", controller.Transfer)
	app.Post("/v1/ewallet/transaction/deposit", controller.Transfer)
	app.Post("/v1/ewallet/transaction/withdrawal", controller.Transfer)
}

func (controller *EWalletController) GetBalance(c *fiber.Ctx) error {
	accountNumberParam := c.Params("account_id")

	err := validation.ValidateAccountNumber("account_id", accountNumberParam)

	if err != nil {
		return exception.ValidationError{Message: err.Error()}
	}

	accountNumber, _ := strconv.Atoi(accountNumberParam)
	response, err := controller.EWalletService.GetEWalletBalance(int32(accountNumber))
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    response,
	})
}

func (controller *EWalletController) GetTransactions(c *fiber.Ctx) error {
	accountNumberParam := c.Params("account_id")

	err := validation.ValidateAccountNumber("account_id", accountNumberParam)

	if err != nil {
		return exception.ValidationError{Message: err.Error()}
	}

	accountNumber, _ := strconv.Atoi(accountNumberParam)
	response, err := controller.EWalletService.GetEWalletTransactions(int32(accountNumber))
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    response,
	})
}

func (controller *EWalletController) Transfer(c *fiber.Ctx) error {
	requestBody := model.TransferRequest{}
	err := c.BodyParser(&requestBody)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	validation.ValidateEWalletTransfer(requestBody)
	err = controller.EWalletService.EWalletTransfer(requestBody.FromAccountNumber, requestBody.ToAccountNumber, requestBody.Amount)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}
	return c.JSON(model.WebResponse{
		Code:    201,
		Status:  "Created",
		Message: "Transfer Succeed",
	})
}
