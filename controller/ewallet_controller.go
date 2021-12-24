package controller

// GET http://localhost:3000/v1/ewallet/balance/{account_id}
// POST http://localhost:3000/v1/ewallet/deposit
// GET http://localhost:3000/v1/ewallet/transactions
// POST http://localhost:3000/v1/ewallet/transfer
// POST http://localhost:3000/v1/ewallet/withdrawal

import (
	"fmt"
	"strconv"
	"time"

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
	app.Post("/account/:from_account_number/transfer", controller.Transfer)

	//app.Get("/v1/ewallet/balance/:account_id", middleware.AuthorizationGuard(controller.GetBalance))
	// api.HandleFunc("/org/customer/finalize", middleware.AuthorizationGuard(handler.Finalization, orgConst.APICodeFinalizationApp, orgConst.MenuCodeEmpty))

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

func (controller *EWalletController) Transfer(c *fiber.Ctx) error {
	fmt.Println("----- Transfer :", time.Now())
	fromAccountNumberParam := c.Params("from_account_number")
	err := validation.ValidateAccountNumber("from_account_number", fromAccountNumberParam)
	if err != nil {
		fmt.Println("----- err 1 :", err)
		return exception.ValidationError{Message: err.Error()}
	}
	fromAccountNumber, _ := strconv.Atoi(fromAccountNumberParam)

	requestBody := model.CreateTransferRequest{}
	err = c.BodyParser(&requestBody)
	if err != nil {
		fmt.Println("----- err 2 :", err)
		return exception.ErrorHandler(c, err)
	}

	validation.ValidateTransfer(requestBody)
	err = controller.EWalletService.EWalletTransfer(int32(fromAccountNumber), requestBody.ToAccountNumber, requestBody.Amount)
	if err != nil {
		fmt.Println("----- err 3 :", err)
		return exception.ErrorHandler(c, err)
	}
	c.Status(201)
	fmt.Println("----- done")
	return nil
}
