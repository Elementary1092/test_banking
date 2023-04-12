package model

import (
	"errors"
	"github.com/Elementary1092/test_banking/internal/entity"
)

var (
	ErrInvalidUpdateParameters  = errors.New("invalid update parameters")
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
)

type UpdateAccount struct {
	to                string
	from              string
	tType             entity.TransactionType
	transactionAmount float64
	currency          string
}

func NewUpdateAccount(to, from, currency string, tType entity.TransactionType, amount float64) (*UpdateAccount, error) {
	if to == "" || from == "" || currency == "" {
		return nil, ErrInvalidUpdateParameters
	}

	if amount < 0 {
		return nil, ErrInvalidTransactionAmount
	}

	return &UpdateAccount{
		to:                to,
		from:              from,
		tType:             tType,
		transactionAmount: amount,
		currency:          currency,
	}, nil
}

func (u *UpdateAccount) To() string {
	return u.to
}

func (u *UpdateAccount) From() string {
	return u.from
}

func (u *UpdateAccount) TransactionType() string {
	return string(u.tType)
}

func (u *UpdateAccount) Currency() string {
	return u.currency
}

func (u *UpdateAccount) Amount() float64 {
	return u.transactionAmount
}
