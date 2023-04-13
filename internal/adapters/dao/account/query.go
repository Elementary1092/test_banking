package account

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	"github.com/Elementary1092/test_banking/internal/domain/account/query/model"
	"strings"
)

type QueryDAO struct {
	db                   dao.DB
	accountTableName     string
	transactionTableName string
}

func NewQueryDAO(db dao.DB) *QueryDAO {
	if db == nil {
		panic("dao DB is nil in dao/account.QueryDAO")
	}

	return &QueryDAO{
		db:                   db,
		accountTableName:     `"accounts"`,
		transactionTableName: `"transactions"`,
	}
}

func (q *QueryDAO) FindAccount(ctx context.Context, params map[string]string) (*model.Account, error) {
	const queryFmt = `SELECT "number", "customer_id", "currency", "balance", "created_at" FROM %s `
	query := fmt.Sprintf(queryFmt, q.accountTableName)

	var whereClause strings.Builder
	for key, value := range params {
		switch key {
		case "account_number":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf("%s = %s", key, value))
		case "user_id":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf("%s = %s", key, value))
		}
	}

	if whereClause.Len() != 0 {
		query = fmt.Sprintf("%s WHERE %s LIMIT 1", query, whereClause.String())
	}

	row := q.db.QueryRow(ctx, query)
	var account accountModel
	if err := row.Scan(
		&account.Number,
		&account.UserID,
		&account.Currency,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		return nil, dao.ResolveError(err)
	}

	response := account.MapToReadModel()
	return response, nil
}

func (q *QueryDAO) FindTransactions(ctx context.Context, accountNumber string, limit, offset uint64) (*model.Account, error) {
	account, err := q.FindAccount(ctx, map[string]string{
		"account_number": accountNumber,
	})
	if err != nil {
		return nil, err
	}

	const queryFmt = `SELECT "id", "from", "to", "type", "currency", "amount", "created_at" 
		FROM %s WHERE "to"=$1 OR "from"=$1`
	query := fmt.Sprintf(
		queryFmt,
		q.transactionTableName,
	)
	if limit != 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	rows, err := q.db.Query(ctx, query, accountNumber)
	if err != nil {
		return nil, dao.ResolveError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction transactionModel
		if err := rows.Scan(
			&transaction.ID,
			&transaction.From,
			&transaction.To,
			&transaction.Type,
			&transaction.Currency,
			&transaction.Amount,
			&transaction.PerformedAt,
		); err != nil {
			return nil, err
		}

		transaction.AddToReadModel(account)
	}

	return account, nil
}

func (q *QueryDAO) ListAccounts(ctx context.Context, params map[string]string, limit, offset uint64) ([]*model.Account, error) {
	const queryFmt = `SELECT "number", "customer_id", "currency", "balance", "created_at" FROM %s `
	query := fmt.Sprintf(queryFmt, q.accountTableName)

	var whereClause strings.Builder
	for key, value := range params {
		switch key {
		case "account_number":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf("%s = %s", key, value))
		case "user_id":
			if whereClause.Len() != 0 {
				whereClause.WriteString(" AND ")
			}
			whereClause.WriteString(fmt.Sprintf("%s = %s", key, value))
		}
	}

	if whereClause.Len() != 0 {
		query = fmt.Sprintf("%s WHERE %s", query, whereClause.String())
	}

	if limit != 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, dao.ResolveError(err)
	}
	defer rows.Close()

	res := make([]*model.Account, 0)
	for rows.Next() {
		var account accountModel
		if err = rows.Scan(
			&account.Number,
			&account.UserID,
			&account.Currency,
			&account.Balance,
			&account.CreatedAt,
		); err != nil {
			return nil, dao.ResolveError(err)
		}

		response := account.MapToReadModel()

		res = append(res, response)
	}

	return res, nil
}
