package auth

import (
    "context"
    "errors"
    "github.com/Elementary1092/test_banking/internal/domain/customer/query"
    "github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
    "github.com/Elementary1092/test_banking/pkg/hasher"
)

var (
    ErrInvalidAuthParameters = errors.New("invalid email or password")
)

type Handler struct {
    repo query.ReadDAO
}

func NewAuthHandler(repo query.ReadDAO) *Handler {
    if repo == nil {
        panic("customer read dao is nil")
    }

    return &Handler{
        repo: repo,
    }
}

func (h *Handler) Handle(ctx context.Context, query Query) (*model.Customer, error) {
    searchParam := map[string]string{
        "email": query.Email,
    }

    customer, err := h.repo.FindCustomer(ctx, searchParam)
    if err != nil {
        return nil, ErrInvalidAuthParameters
    }

    if err := hasher.Verify(customer.Password(), query.Password); err != nil {
        return nil, ErrInvalidAuthParameters
    }

    return customer, nil
}
