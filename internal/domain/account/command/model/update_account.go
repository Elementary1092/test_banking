package model

import (
	"errors"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

var (
	ErrInvalidUpdateParameters  = errors.New("invalid update parameters")
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
)

type UpdateAccount struct {
	at                time.Time
	to                string
	from              string
	tType             entity.TransactionType
	transactionAmount float64
}

func NewUpdateAccount(to, from string, tType entity.TransactionType, amount float64, at time.Time) (*UpdateAccount, error) {
	if to == "" || from == "" {
		return nil, ErrInvalidUpdateParameters
	}

	if amount < 0 {
		return nil, ErrInvalidTransactionAmount
	}

	return &UpdateAccount{
		at:                at,
		to:                to,
		from:              from,
		tType:             tType,
		transactionAmount: amount,
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

func (u *UpdateAccount) Amount() float64 {
	return u.transactionAmount
}
