package orders

import (
	"github.com/devkemc/fundamentos-golang/common"
	"net/http"
)

var (
	errPaymentsIsRequired = common.NewError(http.StatusUnprocessableEntity, "no payments: payments is required")
	errItemsIsRequired    = common.NewError(http.StatusUnprocessableEntity, "no items: items are required")
)
