package validation

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pauluswi/bigben/exception"
	"github.com/pauluswi/bigben/model"
)

func ValidateAccountID(fieldName string, param interface{}) error {
	return validation.Errors{
		fieldName: validation.Validate(&param, validation.Required, is.Int),
	}.Filter()
}

func ValidateEWalletTransfer(request model.TransferRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.FromAccountNumber, validation.Required, validation.Min(1)),
		validation.Field(&request.ToAccountNumber, validation.Required, validation.Min(1)),
		validation.Field(&request.Amount, validation.Required, validation.Min(1)),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	if &request.FromAccountNumber == &request.ToAccountNumber {
		errSameAccount := errors.New("cannot transfer to a same account")
		panic(exception.ValidationError{
			Message: errSameAccount.Error(),
		})
	}
}
