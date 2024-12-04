package orders

import (
	"context"
	"github.com/devkemc/fundamentos-golang/customers"
	"github.com/devkemc/fundamentos-golang/payments"
	"time"
)

type status string

const (
	orderPending   status = "PENDING"
	orderConfirmed status = "CONFIRMED"
	orderCancelled status = "CANCELLED"
)

type Order struct {
	Id         int64               `db:"id" json:"id"`
	Payments   []payments.Payment  `db:"-" json:"payments"`
	Status     status              `db:"status" json:"status"`
	CustomerId int64               `db:"customer_id" json:"customer_id"`
	Items      []Item              `db:"-" json:"items"`
	Customer   *customers.Customer `db:"-" json:"customer;omitempty"`
	CreatedAt  time.Time           `db:"created_at" json:"created_at"`
}

type Item struct {
	ProductId int64 `db:"product_id" json:"product_id"`
	Quantity  int64 `db:"quantity" json:"quantity"`
	OrderId   int64 `db:"order_id" json:"_"`
}

func (o *Order) ValidateToSell(ctx context.Context) error {
	if len(o.Payments) == 0 {
		return errPaymentsIsRequired
	}

	for _, payment := range o.Payments {
		if err := payment.ValidatePayment(); err != nil {
			return err
		}
	}

	if len(o.Items) == 0 {
		return errItemsIsRequired
	}
	//todo: validate items

	return nil
}
