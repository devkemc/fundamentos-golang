package products

import (
	"context"
	"time"
)

type productServiceV1 struct {
	productRepository ProductRepository
}

func (p productServiceV1) GetProductById(ctx context.Context, id int64) (*Product, error) {
	time.Sleep(time.Second * 3)
	return p.productRepository.FindProductById(ctx, id)
}

func NewProductServiceV1(productRepository ProductRepository) ProductService {
	return &productServiceV1{productRepository: productRepository}
}
