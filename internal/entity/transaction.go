package entity

import (
    "errors"
    "time"
)

var (
    ErrInvalidAmount = errors.New("provided amount is invalid")
)

// Transaction is immutable representation of a transaction between customers
type Transaction struct {
    at     time.Time
    amount float64
    from   string
    to     string
}

func NewTransaction(from, to string, amount float64, at time.Time) (*Transaction, error) {
    if amount < 0 {
        return nil, ErrInvalidAmount
    }

    return &Transaction{
        at:     at,
        amount: amount,
        from:   from,
        to:     to,
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
