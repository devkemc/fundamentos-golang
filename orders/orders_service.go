package orders

import "context"

type OrderService interface {
	Sell(ctx context.Context, order *Order) error
	GetOrderDetails(ctx context.Context, orderId int64) (*Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
}
