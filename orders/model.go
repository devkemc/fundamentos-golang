package orders

import "FundamentosGolang/payments"

type Order struct {
	Id         int64              `db:"id" json:"id"`
	Payments   []payments.Payment `db:"-" json:"payments"`
	CustomerId int64              `db:"customer_id" json:"customer_id"`
	Items      []Item             `db:"-" json:"items"`
}

type Item struct {
	ProductId int64 `db:"product_id" json:"product_id"`
	Quantity  int64 `db:"quantity" json:"quantity"`
	OrderId   int64 `db:"order_id" json:"_"`
}
