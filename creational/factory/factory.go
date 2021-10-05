package factory

import "errors"

type PaymentMethod interface {
	Pay(amount float32) string
}

// Our current implemented payment methods are described below
const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	return nil, errors.New("not implemented yet")
}

type CashPM struct{}

type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return ""
}

func (c *DebitCardPM) Pay(amount float32) string {
	return ""
}
