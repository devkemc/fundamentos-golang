package orders

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type orderRepositorySqlx struct {
	common.Repository
}

func NewOrderRepositorySqlx(repository common.Repository) OrderRepository {
	return &orderRepositorySqlx{repository}
}

func (o orderRepositorySqlx) FindOrderById(ctx context.Context, orderId int64) (*Order, error) {
	query := `
		SELECT	id,
				status,
				customer_id,
				created_at
		FROM orders
		WHERE id = ?;
	`
	var order Order
	o.InitTransaction()
	if err := o.GetTx().QueryRowxContext(ctx, query, orderId).StructScan(&order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (o orderRepositorySqlx) GetAllOrders(ctx context.Context) ([]Order, error) {
	query := `
		SELECT	id,
				status,
				customer_id,
				created_at
		FROM orders
	`
	o.InitTransaction()
	defer o.Rollback()
	var orders []Order
	if err := o.GetTx().SelectContext(ctx, &orders, query); err != nil {
		return nil, err
	}
	return orders, nil
}

func (o orderRepositorySqlx) SaveOrder(ctx context.Context, order Order) (int64, error) {
	query := `
	INSERT INTO orders (status, customer_id, created_at)
	VALUES (:status, :customer_id, :created_at)
	`
	args := map[string]interface{}{
		"status":      order.Status,
		"customer_id": order.CustomerId,
		"created_at":  order.CreatedAt,
	}
	tx := o.GetTx()
	result, err := tx.NamedExecContext(ctx, query, args)
	if err != nil {
		return 0, err
	}
	orderId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	query = `
			INSERT INTO items (product_id, quantity, order_id)
			VALUES (:product_id, :quantity, :order_id)
		`
	for _, item := range order.Items {
		args = map[string]interface{}{
			"product_id": item.ProductId,
			"quantity":   item.Quantity,
			"order_id":   orderId,
		}
		_, err := tx.NamedExecContext(ctx, query, args)
		if err != nil {
			return 0, err
		}
	}
	return orderId, nil
}

func (o orderRepositorySqlx) ConfirmOrder(ctx context.Context, orderId int64) error {
	query := `
		UPDATE orders
		SET status = ?
		WHERE id = ?
	`
	args := []interface{}{
		orderConfirmed,
		orderId,
	}
	o.InitTransaction()
	tx := o.GetTx()
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		o.Rollback()
	}
	if err := o.Commit(); err != nil {
		o.Rollback()
		return err
	}

	return nil

}
