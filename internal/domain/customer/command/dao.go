package command

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command/model"
)

type WriteDAO interface {
	CreateCustomer(ctx context.Context, customer *model.WriteModel) error
}
