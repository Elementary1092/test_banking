package transactions

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/query"
	"github.com/Elementary1092/test_banking/internal/domain/account/query/model"
)

type Handler struct {
	repo query.ReadDAO
}

func NewHandler(repo query.ReadDAO) *Handler {
	if repo == nil {
		panic("query ReadDAO is nil in account.query.transactions")
	}

	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, qry Query) (*model.Account, error) {
	return h.repo.FindTransactions(ctx, qry.AccountNumber, qry.Limit, qry.Offset)
}
