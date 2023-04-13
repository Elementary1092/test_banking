package replenish

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/command"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

type Handler struct {
	repo command.WriteDAO
}

func NewHandler(repo command.WriteDAO) *Handler {
	if repo == nil {
		panic("command write dao is nil in replenish.Handler")
	}

	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	if cmd.Amount < 0 {
		return errResponses.ErrInvalidTransactionAmount
	}

	toAccount, err := h.repo.FindAccount(ctx, map[string]string{
		"account_number": cmd.AccountNumber,
	})
	if err != nil {
		return err
	}

	// Logically, may be, currency
	// should have been taken from command and compared with account's currency,
	// but for this app I think it is not critical.
	updateModel, err := model.NewUpdateAccount(
		cmd.AccountNumber,
		cmd.FromCard,
		toAccount.Currency(),
		entity.ReplenishType,
		cmd.Amount,
		time.Now())
	if err != nil {
		return err
	}

	return h.repo.AddTransaction(ctx, updateModel, entity.ToAccount)
}
