package orders

import (
	"context"
	"github.com/devkemc/fundamentos-golang/customers"
	"github.com/devkemc/fundamentos-golang/emails"
	"github.com/devkemc/fundamentos-golang/payments"
	"github.com/devkemc/fundamentos-golang/products"
)

type orderServiceV1 struct {
	orderRepository OrderRepository
	emailService    emails.EmailService
	paymentService  payments.PaymentService
	customerService customers.CustomerService
	productService  products.ProductService
}

func (o orderServiceV1) GetOrderDetails(ctx context.Context, orderId int64) (*Order, error) {
	order, err := o.orderRepository.FindOrderById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	order.Payments, err = o.paymentService.GetPaymentsByOrderId(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	order.Customer, err = o.customerService.GetCustomerById(ctx, order.CustomerId)
	if err != nil {
		return nil, err
	}

	order.Items, err = o.orderRepository.FindItemsByOrderId(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	for idx := range order.Items {
		item := &order.Items[idx]
		item.Product, err = o.productService.GetProductById(ctx, item.ProductId)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (o orderServiceV1) GetOrders(ctx context.Context) ([]Order, error) {
	return o.orderRepository.GetAllOrders(ctx)
}

func (o orderServiceV1) Sell(ctx context.Context, order *Order) error {
	if err := order.ValidateToSell(ctx); err != nil {
		return err
	}

	err := o.calculateAmount(ctx, order)
	if err != nil {
		return err
	}

	o.orderRepository.InitTransaction(ctx)

	order.Status = orderPending
	orderId, err := o.orderRepository.SaveOrder(ctx, *order)
	if err != nil {
		o.orderRepository.Rollback()
		return err
	}

	err = o.orderRepository.Commit()
	if err != nil {
		return err
	}

	order.Id = orderId
	if err := o.paymentService.ProcessPayments(ctx, order.Payments, orderId); err != nil {
		return err
	}

	if err := o.orderRepository.ConfirmOrder(ctx, order.Id); err != nil {
		return err
	}

	if err := o.emailService.SendEmail(ctx, emails.Email{}); err != nil {
		return err
	}

	return nil
}

func (o orderServiceV1) calculateAmount(ctx context.Context, order *Order) error {
	for idx := range order.Items {
		item := &order.Items[idx]
		product, err := o.productService.GetProductById(ctx, item.ProductId)
		if err != nil {
			return err
		}
		item.Amount = float64(item.Quantity) * product.Price
	}
	return nil
}

func NewOrderServiceV1(orderRepository OrderRepository, emailService emails.EmailService, paymentService payments.PaymentService, customerService customers.CustomerService, productService products.ProductService) OrderService {
	return &orderServiceV1{orderRepository, emailService, paymentService, customerService, productService}
}
