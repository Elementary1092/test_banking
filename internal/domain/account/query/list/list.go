package list

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/account/query"
	"github.com/Elementary1092/test_banking/internal/domain/account/query/model"
)

type Handler struct {
	dao query.ReadDAO
}

func NewHandler(dao query.ReadDAO) *Handler {
	if dao == nil {
		panic("read DAO is nil in account/query/list.Handler")
	}

	return &Handler{
		dao: dao,
	}
}

func (h *Handler) Handle(ctx context.Context, qry Query) ([]*model.Account, error) {
	params := make(map[string]string)

	if qry.UserID != "" {
		params["user_id"] = qry.UserID
	}

	return h.dao.ListAccounts(ctx, params, qry.Limit, qry.Offset)
}
