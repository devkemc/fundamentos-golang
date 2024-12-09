package customers

import (
	"context"
	"github.com/devkemc/fundamentos-golang/common"
)

type CustomerRepositorySqlx struct {
	common.Repository
}

func (c CustomerRepositorySqlx) FindCustomerById(ctx context.Context, customerId int) (*Customer, error) {
	query := `
		SELECT	id,
		    	name,
		    	email
		FROM customers
		WHERE id = ?;
	`
	args := []interface{}{customerId}
	var customer Customer
	if err := c.GetTx(ctx).QueryRowxContext(ctx, query, args...).StructScan(&customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

func NewCustomerRepositorySqlx(commonRepo common.Repository) CustomerRepository {
	return &CustomerRepositorySqlx{Repository: commonRepo}
}
