// No need to use here CQRS pattern

package token

import (
	"context"
	"github.com/Elementary1092/test_banking/internal/entity"
)

type DAO interface {
	Insert(ctx context.Context, token *entity.RefreshToken) error
	Delete(ctx context.Context, token string) error
	Exists(ctx context.Context, token string) (bool, error)
}
