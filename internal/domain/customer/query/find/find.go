package find

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
)

type Handler struct {
	repo query.ReadDAO
}

func NewFindHandler(repo query.ReadDAO) *Handler {
	if repo == nil {
		panic("customer read dao is nil")
	}

	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, query Query) (*model.Customer, error) {
	customer, err := h.repo.FindCustomer(ctx, query.Params)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
