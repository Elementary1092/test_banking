package replenish

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/domain/account/command"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

var (
	ErrInvalidReplenishAmount = errors.New("replenish amount cannot be less than 0")
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
		return ErrInvalidReplenishAmount
	}

	updateModel, err := model.NewUpdateAccount(
		cmd.AccountNumber,
		cmd.FromCard,
		entity.ReplenishType,
		cmd.Amount,
		time.Now())
	if err != nil {
		return err
	}

	return h.repo.UpdateAccount(ctx, updateModel, entity.ToAccount)
}
