package command

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
)

type UpdateType uint32

const (
	UpdateOnlyToAccount UpdateType = iota
	UpdateOnlyFromAccount
	UpdateBothAccounts
)

type WriteDAO interface {
	// FindAccount should be able to search by 2 parameters: "account_number" and "user_id"
	FindAccount(ctx context.Context, params map[string]string) (*model.Account, error)
	CreateAccount(ctx context.Context, account *model.Account) error

	// UpdateAccount should make update according to update type in 1 transaction
	UpdateAccount(ctx context.Context, updateReq *model.UpdateAccount, t UpdateType) error
}
