package payments

import (
	"context"
	"time"
)

type paymentsServiceSimulator struct {
	paymentRepository PaymentRepository
}

func (p paymentsServiceSimulator) ProcessPayments(ctx context.Context, payments []Payment, orderId int64) error {
	for _, payment := range payments {
		time.Sleep(time.Second * 10)
		payment.OrderId = orderId
		_, err := p.paymentRepository.SavePayment(ctx, payment)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p paymentsServiceSimulator) GetPaymentsByOrderId(ctx context.Context, orderId int64) ([]Payment, error) {
	return p.paymentRepository.FindPaymentsByOrderId(ctx, orderId)
}

func NewPaymentsServiceSimulator(paymentRepo PaymentRepository) PaymentService {
	return &paymentsServiceSimulator{paymentRepository: paymentRepo}
}
