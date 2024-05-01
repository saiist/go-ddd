package valueobject

import "errors"

type Money struct {
	Amount   float64
	Currency string
}

// NewMoney creates a new Money object
func NewMoney(amount float64, currency string) (*Money, error) {
	if currency == "" {
		return nil, errors.New("currency cannot be empty")
	}

	return &Money{Amount: amount, Currency: currency}, nil
}

// Add adds the amount of the given Money object to the current Money object
func (m *Money) Add(arg Money) error {
	if arg.Currency != m.Currency {
		return errors.New("currency mismatch")
	}

	m.Amount += arg.Amount
	arg.Amount = 0
	return nil
}
