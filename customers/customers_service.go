package customers

import "context"

type CustomerService interface {
	GetCustomerById(ctx context.Context, id int64) (*Customer, error)
}
