package controller

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

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService *service.AccountService) AccountController {
	return AccountController{
		AccountService: *accountService,
	}
}

func (controller *AccountController) Route(app *fiber.App) {
	app.Get("/account/:account_number", controller.GetBalance)
	app.Post("/account/:from_account_number/transfer", controller.Transfer)
}

func (controller *AccountController) GetBalance(c *fiber.Ctx) error {
	accountNumberParam := c.Params("account_number")

	err := validation.ValidateAccountNumber("account_number", accountNumberParam)

	if err != nil {
		return exception.ValidationError{Message: err.Error()}
	}

	accountNumber, _ := strconv.Atoi(accountNumberParam)
	response, err := controller.AccountService.GetBalance(int32(accountNumber))
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

func (controller *AccountController) Transfer(c *fiber.Ctx) error {
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
	err = controller.AccountService.Transfer(int32(fromAccountNumber), requestBody.ToAccountNumber, requestBody.Amount)
	if err != nil {
		fmt.Println("----- err 3 :", err)
		return exception.ErrorHandler(c, err)
	}
	c.Status(201)
	fmt.Println("----- done")
	return nil
}
