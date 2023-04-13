package customer

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"strings"
)

type QueryDAO struct {
	db                dao.DB
	customerTableName string
}

func NewQueryDAO(db dao.DB) *QueryDAO {
	if db == nil {
		panic("db is nil in dao/customer.QueryDAO")
	}

	return &QueryDAO{
		db:                db,
		customerTableName: `"customers"`,
	}
}

func (q *QueryDAO) FindCustomer(ctx context.Context, filter map[string]string) (*model.Customer, error) {
	const queryFmt = `SELECT "uuid", "email", "password", "created_at" FROM %s`
	query := fmt.Sprintf(queryFmt, q.customerTableName)

	var whereClause strings.Builder
	var params = make([]any, 0)
	for key, value := range filter {
		switch key {
		case "uuid":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf(`"%s" = $%d`, key, len(params)+1))
			params = append(params, value)
		case "email":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf(`"%s" = $%d`, key, len(params)+1))
			params = append(params, value)
		}
	}

	if whereClause.Len() != 0 {
		query = fmt.Sprintf("%s WHERE %s", query, whereClause.String())
	}

	var customer customerModel
	if err := q.db.QueryRow(ctx, query, params...).Scan(
		&customer.UUID,
		&customer.Email,
		&customer.Password,
		&customer.CreatedAt,
	); err != nil {
		return nil, dao.ResolveError(err)
	}

	return customer.MapToReadModel(), nil
}
