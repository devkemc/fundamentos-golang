package payments

import "slices"

type paymentType string

const (
	typeCredit paymentType = "CREDIT"
)

type paymentStatus string

const (
	pStatusPending  paymentStatus = "PENDING"
	pStatusRejected paymentStatus = "REJECTED"
	pStatusCanceled paymentStatus = "CANCELED"
	pStatusFailed   paymentStatus = "FAILED"
	pStatusAccepted paymentStatus = "ACCEPTED"
)

var validPaymentTypes = []paymentType{typeCredit}

type Payment struct {
	Id      int64         `json:"id" db:"id"`
	Amount  float32       `json:"amount" db:"amount"`
	Type    paymentType   `json:"type" db:"type"`
	Status  paymentStatus `json:"status" db:"status"`
	OrderId int64         `json:"order" db:"order_id"`
}

func (p *Payment) ValidatePayment() error {
	if !slices.Contains(validPaymentTypes, p.Type) {
		return errPaymentTypeIsInvalid
	}
	if p.Amount < 0 {
		return errPaymentAmountIsInvalid
	}
	return nil
}
