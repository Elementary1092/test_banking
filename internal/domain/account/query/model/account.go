package model

import "github.com/Elementary1092/test_banking/internal/entity"

type Account struct {
	Info *entity.Account

	// account should belong to some registered user
	UserID string

	// account may have many outgoing transactions
	Transactions []*entity.Transaction
}
