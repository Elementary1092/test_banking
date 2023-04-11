package query

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
)

type ReadDAO interface {
	FindCustomer(ctx context.Context, filter map[string]string) (*model.Customer, error)
}
