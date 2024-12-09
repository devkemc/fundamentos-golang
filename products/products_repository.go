package products

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type ProductRepository interface {
	common.Repository
	FindProductById(ctx context.Context, id int64) (*Product, error)
}
