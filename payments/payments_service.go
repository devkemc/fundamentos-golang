package payments

import "context"

type PaymentService interface {
	ProcessPayments(ctx context.Context, payments []Payment, orderId int64) error
	GetPaymentsByOrderId(ctx context.Context, orderId int64) ([]Payment, error)
}
