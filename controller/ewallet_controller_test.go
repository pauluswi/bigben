package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/pauluswi/bigben/model"
	"github.com/stretchr/testify/assert"
)

func TestEWalletController_GetBalance(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/ewallet/balance/10001", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
}

func TestEWalletController_GetTransactionHistory(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/ewallet/transaction/history/10001", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
}

func TestEWalletController_Transfer(t *testing.T) {
	makeTransfer := model.EWalletTransferRequest{
		FromAccountNumber: 10001,
		ToAccountNumber:   10002,
		Amount:            100,
	}

	requestBody, _ := json.Marshal(makeTransfer)

	request := httptest.NewRequest("POST", "/v1/ewallet/transaction/transfer", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
}

func TestEWalletController_Deposit(t *testing.T) {
	makeTransfer := model.EWalletDepositRequest{
		ToAccountNumber: 10002,
		Amount:          10,
	}

	requestBody, _ := json.Marshal(makeTransfer)

	request := httptest.NewRequest("POST", "/v1/ewallet/transaction/deposit", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
}

func TestEWalletController_Withdrawal(t *testing.T) {
	makeTransfer := model.EWalletWitdrawalRequest{
		FromAccountNumber: 10001,
		Amount:            10,
	}

	requestBody, _ := json.Marshal(makeTransfer)

	request := httptest.NewRequest("POST", "/v1/ewallet/transaction/withdrawal", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
}

func TestEWalletController_HealthCheck(t *testing.T) {
	request := httptest.NewRequest("GET", "/ping", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
}
