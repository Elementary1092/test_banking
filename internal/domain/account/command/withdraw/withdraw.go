package withdraw

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/domain/account/command"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient balance")
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
		return ErrInvalidAmount
	}
	fromAccount, err := h.repo.FindAccount(ctx, map[string]string{
		"account_number": cmd.From,
	})
	if err != nil {
		return err
	}

	if fromAccount.Balance() < cmd.Amount {
		return ErrInsufficientFunds
	}

	updatedFrom, err := model.NewUpdateAccount(cmd.To, cmd.From, entity.WithdrawType, cmd.Amount, time.Now())
	if err != nil {
		return err
	}

	return h.repo.UpdateAccount(ctx, updatedFrom, command.UpdateOnlyFromAccount)
}
