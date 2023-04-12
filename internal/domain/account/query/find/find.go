package find

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
        panic("read DAO is nil in account.query.find")
    }

    return &Handler{
        repo: repo,
    }
}

func (h *Handler) Handle(ctx context.Context, qry Query) (*model.Account, error) {
    params := make(map[string]string, 0)

    if qry.AccountNumber != "" {
        params["account_number"] = qry.AccountNumber
    }

    if qry.UserID != "" {
        params["user_id"] = qry.UserID
    }

    return h.repo.FindAccount(ctx, params)
}
