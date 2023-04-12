package model

import (
	"github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
	"github.com/Elementary1092/test_banking/internal/entity"
)

type Account struct {
	account *entity.Account

	userID string

	// account write model may work only with 1 transaction at a time
	transaction *entity.Transaction
}

func NewWriteAccount(accountNumber, currency, userID string, balance float64, transaction *entity.Transaction) (*Account, error) {
	if accountNumber == "" || currency == "" || userID == "" {
		return nil, errResponses.ErrInvalidAccountInfo
	}
	if balance < 0 {
		return nil, errResponses.ErrInvalidBalanceInfo
	}

	return &Account{
		account: &entity.Account{
			Number:   accountNumber,
			Currency: currency,
			Balance:  balance,
		},
		userID:      userID,
		transaction: transaction,
	}, nil
}

func (a *Account) Number() string {
	return a.account.Number
}

func (a *Account) Currency() string {
	return a.account.Currency
}

func (a *Account) UserID() string {
	return a.userID
}

func (a *Account) Balance() float64 {
	return a.account.Balance
}

func (a *Account) Transaction() *entity.Transaction {
	return a.transaction
}
