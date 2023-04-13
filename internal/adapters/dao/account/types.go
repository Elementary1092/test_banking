package account

import (
	writeModel "github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	readModel "github.com/Elementary1092/test_banking/internal/domain/account/query/model"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

type accountModel struct {
	CreatedAt time.Time
	Number    string
	UserID    string
	Currency  string
	Balance   float64
}

func (a accountModel) MapToReadModel() *readModel.Account {
	accountInfo := entity.Account{
		Number:   a.Number,
		Currency: a.Currency,
		Balance:  a.Balance,
	}

	account := readModel.Account{
		Info:   &accountInfo,
		UserID: a.UserID,
	}

	return &account
}

func (a accountModel) MapToWriteModel() (*writeModel.Account, error) {
	account, err := writeModel.NewWriteAccount(a.Number, a.Currency, a.UserID, a.Balance, nil, a.CreatedAt)
	return account, err
}

type transactionModel struct {
	PerformedAt time.Time
	From        string
	To          string
	Type        string
	Currency    string
	Amount      float64
	ID          uint64
}

func (t transactionModel) AddToReadModel(account *readModel.Account) {
	// Assuming that all transaction are written correctly
	// and this function should not return an error.
	transaction, _ := entity.NewTransaction(
		t.From,
		t.To,
		t.Currency,
		entity.TransactionType(t.Type),
		t.Amount,
		t.PerformedAt,
	)
	account.Transactions = append(account.Transactions, transaction)
}
