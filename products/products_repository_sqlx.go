package products

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type productRepositorySqlx struct {
	common.Repository
}

func (p productRepositorySqlx) FindProductById(ctx context.Context, id int64) (*Product, error) {
	query := `
		SELECT	id,
		    	name,
		    	description,
		    	price
		FROM	products
		WHERE id = ?
	`

	args := []interface{}{id}
	var product Product
	if err := p.GetTx(ctx).QueryRowxContext(ctx, query, args...).StructScan(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func NewProductRepositorySqlx(repository common.Repository) ProductRepository {
	return &productRepositorySqlx{
		Repository: repository,
	}
}
