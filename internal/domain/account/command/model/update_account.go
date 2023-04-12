package model

import (
	"github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
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
		return nil, errResponses.ErrInvalidUpdateParameters
	}

	if amount < 0 {
		return nil, errResponses.ErrInvalidTransactionAmount
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
