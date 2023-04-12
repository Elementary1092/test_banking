package command

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
)

type WriteDAO interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	UpdateBalance(ctx context.Context, accountNumber string, balance float64) error
	Exists(ctx context.Context, accountNumber string) error
	AddAccountTransaction(ctx context.Context, account *model.Account) error
}
