package payments

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type paymentRepositorySqlx struct {
	common.Repository
}

func (p paymentRepositorySqlx) SavePayment(ctx context.Context, payment Payment) (int64, error) {
	query := `
		INSERT INTO payments (amount, type, status, order_id)
		VALUES (:amount, :type, :status, :order_id)
	`
	args := map[string]interface{}{
		"amount":   payment.Amount,
		"type":     payment.Type,
		"status":   payment.Status,
		"order_id": payment.OrderId,
	}
	result, err := p.GetTx(ctx).NamedExecContext(ctx, query, args)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p paymentRepositorySqlx) FindPaymentsByOrderId(ctx context.Context, orderId int64) ([]Payment, error) {
	query := `
	SELECT id, amount, type, status, order_id
	FROM payments
	WHERE order_id = ?
	`
	var payments []Payment

	defer p.Rollback()
	if err := p.GetTx(ctx).SelectContext(ctx, &payments, query, orderId); err != nil {
		return nil, err
	}
	return payments, nil
}

func NewPaymentRepositorySqlx(repo common.Repository) PaymentRepository {
	return &paymentRepositorySqlx{repo}
}
