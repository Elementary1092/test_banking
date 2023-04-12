package model

import "github.com/Elementary1092/test_banking/internal/entity"

type Account struct {
    Info *entity.Account

    // account should belong to some registered user
    UserID string

    // account may have many outgoing transactions
    outTransactions []*entity.Transaction

    // account may be addressed by many transactions
    inTransactions []*entity.Transaction
}

func (a *Account) OutTransactions() []*entity.Transaction {
    return a.outTransactions
}

func (a *Account) AddOutTransactions(transactions ...*entity.Transaction) {
    if len(transactions) == 0 {
        return
    }

    a.outTransactions = append(a.outTransactions, transactions...)
}

func (a *Account) InTransactions() []*entity.Transaction {
    return a.inTransactions
}

func (a *Account) AddInTransactions(transactions ...*entity.Transaction) {
    if len(transactions) == 0 {
        return
    }

    a.outTransactions = append(a.outTransactions, transactions...)
}
