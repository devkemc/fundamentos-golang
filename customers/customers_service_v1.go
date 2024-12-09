package customers

import "context"

type customerServiceV1 struct {
	customerRepository CustomerRepository
}

func (c customerServiceV1) GetCustomerById(ctx context.Context, id int) (*Customer, error) {
	return c.customerRepository.FindCustomerById(ctx, id)
}

func NewCustomerServiceV1(customerRepository CustomerRepository) CustomerService {
	return &customerServiceV1{customerRepository: customerRepository}
}
