package orders

import (
	"context"
	"fmt"
	"github.com/devkemc/fundamentos-golang/customers"
	"github.com/devkemc/fundamentos-golang/emails"
	"github.com/devkemc/fundamentos-golang/payments"
	"github.com/devkemc/fundamentos-golang/products"
	"github.com/gofiber/fiber/v2/log"
	"sync"
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

	errorGroupChan := make(chan error, 3)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		order.Payments, err = o.paymentService.GetPaymentsByOrderId(ctx, order.Id)
		if err != nil {
			errorGroupChan <- err
		}
	}(ctx, &wg)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		order.Customer, err = o.customerService.GetCustomerById(ctx, order.CustomerId)
		if err != nil {
			errorGroupChan <- err
		}
	}(ctx, &wg)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		order.Items, err = o.orderRepository.FindItemsByOrderId(ctx, order.Id)
		if err != nil {
			errorGroupChan <- err
		}
		items := make(map[int64]*Item, len(order.Items))

		productsChan := make(chan map[int64]*products.Product, len(order.Items))
		productsChanErr := make(chan error, len(order.Items))
		wgProducts := sync.WaitGroup{}

		for idx := range order.Items {
			item := &order.Items[idx]

			items[item.Id] = item

			go func(ctx context.Context, wg *sync.WaitGroup) {
				defer wg.Done()
				itemId := item.Id
				product, errProduct := o.productService.GetProductById(ctx, item.ProductId)
				productsChan <- map[int64]*products.Product{
					itemId: product,
				}
				productsChanErr <- errProduct
			}(ctx, &wgProducts)
		}

		wgProducts.Wait()

		close(productsChan)
		close(productsChanErr)

		for errProduct := range productsChanErr {
			if errProduct != nil {
				err = fmt.Errorf("%w %w", err, errProduct)
			}
		}

		for productMap := range productsChan {
			for itemId, product := range productMap {
				items[itemId].Product = product
			}
		}

		if err != nil {
			errorGroupChan <- err
		}

	}(ctx, &wg)

	wg.Wait()
	close(errorGroupChan)
	for errGroup := range errorGroupChan {
		if errGroup != nil {
			err = fmt.Errorf("%w %w", err, errGroup)
			return nil, err
		}
	}

	if err != nil {
		return nil, err
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

	go func() {
		if err := o.paymentService.ProcessPayments(ctx, order.Payments, orderId); err != nil {
			log.Error("Error processing payments: ", err)
		}

		if err := o.orderRepository.ConfirmOrder(ctx, order.Id); err != nil {
			log.Error("Error processing payments: ", err)
		}

		if err := o.emailService.SendEmail(ctx, emails.Email{}); err != nil {
			log.Error("Error processing payments: ", err)
		}
	}()

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
