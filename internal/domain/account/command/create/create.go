package create

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/command"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
)

type Handler struct {
	repo command.WriteDAO
}

func NewCreateHandler(repo command.WriteDAO) *Handler {
	if repo == nil {
		panic("account write dao is nil in command/create")
	}

	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	accountNumber, err := generateAccountNumber(cmd.Currency)
	if err != nil {
		return errResponses.NewInternal("account/command/create.Handle", err)
	}

	account, err := model.NewWriteAccount(accountNumber, cmd.Currency, cmd.UserID, 0, nil)
	if err != nil {
		return err
	}

	return h.repo.CreateAccount(ctx, account)
}
