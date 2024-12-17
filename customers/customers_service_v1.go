package customers

import (
	"context"
	"time"
)

type customerServiceV1 struct {
	customerRepository CustomerRepository
}

func (c customerServiceV1) GetCustomerById(ctx context.Context, id int64) (*Customer, error) {
	time.Sleep(time.Second * 10)
	return c.customerRepository.FindCustomerById(ctx, id)
}

func NewCustomerServiceV1(customerRepository CustomerRepository) CustomerService {
	return &customerServiceV1{customerRepository: customerRepository}
}
