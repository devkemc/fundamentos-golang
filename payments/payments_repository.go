package payments

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type PaymentRepository interface {
	common.Repository
	SavePayment(ctx context.Context, payment Payment) (int64, error)
	FindPaymentsByOrderId(ctx context.Context, orderId int64) ([]Payment, error)
}
