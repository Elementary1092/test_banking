package customer

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command/model"
)

type CommandDAO struct {
	db                 dao.DB
	customersTableName string
}

func NewCommandDAO(db dao.DB) *CommandDAO {
	if db == nil {
		panic("db is nil in dao/customer.CommandDAO")
	}

	return &CommandDAO{
		db:                 db,
		customersTableName: `"customers"`,
	}
}

func (c *CommandDAO) CreateCustomer(ctx context.Context, customer *model.WriteModel) error {
	const queryFmt = `INSERT INTO %s ("uuid", "email", "password", "created_at") 
        VALUES ($1, $2, $3, $4)`

	query := fmt.Sprintf(queryFmt, c.customersTableName)

	if _, err := c.db.Exec(
		ctx,
		query,
		customer.UserID(),
		customer.Email(),
		customer.Password(),
		customer.CreatedAt(),
	); err != nil {
		return dao.ResolveError(err)
	}

	return nil
}
