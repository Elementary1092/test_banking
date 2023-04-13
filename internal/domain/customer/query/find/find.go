package find

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"strings"
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
	params := make(map[string]string)
	if query.UUID = strings.TrimSpace(query.UUID); query.UUID != "" {
		params["uuid"] = query.UUID
	}

	if query.Email = strings.TrimSpace(query.Email); query.Email != "" {
		params["email"] = query.Email
	}

	customer, err := h.repo.FindCustomer(ctx, params)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
