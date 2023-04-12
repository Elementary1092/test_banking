package query

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/query/model"
)

type ReadDAO interface {
	// FindAccount should be able to search by 2 parameters: "account_number" and "user_id"
	FindAccount(ctx context.Context, params map[string]string) (*model.Account, error)
	FindOutTransactions(ctx context.Context, accountNumber string) (*model.Account, error)
	FindInTransactions(ctx context.Context, accountNumber string) (*model.Account, error)
	FindAllTransactions(ctx context.Context, accountNumber string) (*model.Account, error)
}
