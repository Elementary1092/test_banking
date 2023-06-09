package transfer

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
		panic("command write dao is nil in transfer.Handler")
	}

	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	if cmd.Amount < 0 {
		return errResponses.ErrInvalidTransactionAmount
	}
	fromAccount, err := h.repo.FindAccount(ctx, map[string]string{
		"user_id":        cmd.UserID,
		"account_number": cmd.From,
	})
	if err != nil {
		return err
	}

	toAccount, err := h.repo.FindAccount(ctx, map[string]string{
		"account_number": cmd.To,
	})
	if err != nil {
		return err
	}

	if toAccount.Currency() != fromAccount.Currency() {
		return errResponses.ErrInvalidRecipientCurrency
	}

	if fromAccount.Balance() < cmd.Amount {
		return errResponses.ErrInsufficientFunds
	}

	updatedFrom, err := model.NewUpdateAccount(
		cmd.To,
		cmd.From,
		toAccount.Currency(),
		entity.TransferType,
		cmd.Amount,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return h.repo.AddTransaction(ctx, updatedFrom, entity.BothAccounts)
}
