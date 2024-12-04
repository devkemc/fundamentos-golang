package payments

import (
	"github.com/devkemc/fundamentos-golang/common"
	"net/http"
)

var (
	errPaymentTypeIsInvalid   = common.NewError(http.StatusUnprocessableEntity, "payment type is invalid")
	errPaymentAmountIsInvalid = common.NewError(http.StatusUnprocessableEntity, "amount is invalid, amount must be greater than zero")
)
