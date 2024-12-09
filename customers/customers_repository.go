package customers

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type CustomerRepository interface {
	common.Repository
	FindCustomerById(ctx context.Context, customerId int) (*Customer, error)
}
