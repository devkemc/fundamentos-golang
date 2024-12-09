package products

import "context"

type productServiceV1 struct {
	productRepository ProductRepository
}

func (p productServiceV1) GetProductById(ctx context.Context, id int64) (*Product, error) {
	return p.productRepository.FindProductById(ctx, id)
}

func NewProductServiceV1(productRepository ProductRepository) ProductService {
	return &productServiceV1{productRepository: productRepository}
}
