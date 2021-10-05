package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

// Our current implemented payment methods are described below
const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
		case Cash:
			return new(CashPM), nil
		case DebitCard:
			return new(DebitCardPM), nil
		default:
			return nil, errors.New(fmt.Sprintf("Payment Method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
