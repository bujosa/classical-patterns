package factory

type PaymentMethod interface {
	Pay(amount float32) string
}

// Our current implemented payment methods are described below
const (
	Cash      = 1
	DebitCard = 2
)