package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidAmount = errors.New("provided amount is invalid")
)

type TransactionType string

const (
	ReplenishType TransactionType = "replenish"
	WithdrawType  TransactionType = "withdraw"
	TransferType  TransactionType = "transfer"
)

// Transaction is immutable representation of a transaction between customers
type Transaction struct {
	at       time.Time
	currency string
	from     string
	to       string
	tType    TransactionType // transaction type may be 'replenish', 'withdraw', 'transfer'
	amount   float64
}

func NewTransaction(from, to, currency string, tType TransactionType, amount float64, at time.Time) (*Transaction, error) {
	if amount < 0 {
		return nil, ErrInvalidAmount
	}

	return &Transaction{
		at:       at,
		currency: currency,
		amount:   amount,
		from:     from,
		to:       to,
		tType:    tType,
	}, nil
}

func (t *Transaction) From() string {
	return t.from
}

func (t *Transaction) To() string {
	return t.to
}

func (t *Transaction) PerformedAt() time.Time {
	return t.at
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Currency() string {
	return t.currency
}

func (t *Transaction) Type() string {
	return string(t.tType)
}
