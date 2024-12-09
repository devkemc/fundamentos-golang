package orders

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type OrderRepository interface {
	common.Repository
	SaveOrder(ctx context.Context, order Order) (int64, error)
	ConfirmOrder(ctx context.Context, orderId int64) error
	FindOrderById(ctx context.Context, orderId int64) (*Order, error)
	FindItemsByOrderId(ctx context.Context, orderId int64) ([]Item, error)
	GetAllOrders(ctx context.Context) ([]Order, error)
}
