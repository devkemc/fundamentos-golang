package products

import "context"

type ProductService interface {
	GetProductById(ctx context.Context, id int64) (*Product, error)
}
