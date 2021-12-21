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

func TestAccountController_GetBalance(t *testing.T) {
	request := httptest.NewRequest("GET", "/account/55501", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
}

func TestAccountController_Transfer(t *testing.T) {
	makeTransfer := model.CreateTransferRequest{
		ToAccountNumber: 55502,
		Amount:          100,
	}

	requestBody, _ := json.Marshal(makeTransfer)

	request := httptest.NewRequest("POST", "/account/55501/transfer", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	// request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
}
