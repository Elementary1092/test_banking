package account

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	"github.com/Elementary1092/test_banking/internal/entity"
	"strings"
)

type CommandDAO struct {
	db                   dao.DB
	accountTableName     string
	transactionTableName string
}

func NewCommandDAO(db dao.DB) *CommandDAO {
	if db == nil {
		panic("dao DB is nil in dao/account.CommandDAO")
	}

	return &CommandDAO{
		db:                   db,
		accountTableName:     `"accounts"`,
		transactionTableName: `"transactions"`,
	}
}

func (c *CommandDAO) FindAccount(ctx context.Context, params map[string]string) (*model.Account, error) {
	// This function is duplicate of query/QueryDAO.FindAccount.
	// Duplication is added to follow CQRS pattern
	// where read and write operation should be implemented separately.
	// If in the future, read and write db will be separated it will be easier to implement
	// due to this design decision
	const queryFmt = `SELECT "number", "customer_id", "currency", "balance", "created_at" FROM %s `
	query := fmt.Sprintf(queryFmt, c.accountTableName)

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

	row := c.db.QueryRow(ctx, query)
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

	response, err := account.MapToWriteModel()
	return response, err
}

func (c *CommandDAO) CreateAccount(ctx context.Context, account *model.Account) error {
	const queryFmt = `INSERT INTO %s ("number", "customer_id", "currency", "balance", "created_at")
		VALUES ($1, $2, $3, $4, $5)`

	query := fmt.Sprintf(queryFmt, c.accountTableName)

	_, err := c.db.Exec(
		ctx,
		query,
		account.Number(),
		account.UserID(),
		account.Currency(),
		account.Balance(),
		account.CreatedAt(),
	)
	if err != nil {
		return dao.ResolveError(err)
	}

	return nil
}

func (c *CommandDAO) AddTransaction(ctx context.Context, updateReq *model.UpdateAccount, t entity.AppAccount) error {
	// Not a good design decision to transfer logic of balance increase and decrease to dao
	// but did not figure out other method to do this in one transaction without violating CQRS pattern
	const increaseBalanceFmt = `UPDATE %s SET "balance" = "balance" + $2 WHERE "number" = $1`
	const decreaseBalanceFmt = `UPDATE %s SET "balance" = "balance" - $2 WHERE "number" = $1`
	const addTransactionFmt = `INSERT INTO %s ("from", "to", "type", "currency", "amount", "created_at") 
        VALUES ($1, $2, $3, $4, $5, $6)`

	increaseBalance := fmt.Sprintf(increaseBalanceFmt, c.accountTableName)
	decreaseBalance := fmt.Sprintf(decreaseBalanceFmt, c.accountTableName)
	addTransaction := fmt.Sprintf(addTransactionFmt, c.transactionTableName)

	tx, err := c.db.Begin(ctx)
	if err != nil {
		return dao.ResolveError(err)
	}
	// According to the documentation it is safe to call even if transaction will be committed successfully.
	// Also, assuming that rollback will not fail.
	defer tx.Rollback(ctx)

	// Here result does not matter
	_, err = tx.Exec(
		ctx,
		addTransaction,
		updateReq.From(),
		updateReq.To(),
		updateReq.TransactionType(),
		updateReq.Currency(),
		updateReq.Amount(),
		updateReq.At(),
	)
	if err != nil {
		return dao.ResolveError(err)
	}

	if t == entity.ToAccount || t == entity.BothAccounts {
		if res, err := tx.Exec(ctx, increaseBalance, updateReq.To(), updateReq.Amount()); err != nil {
			return dao.ResolveError(err)
		} else if res.RowsAffected() == 0 {
			return dao.NewNotFound(updateReq.To())
		}
	}

	if t == entity.FromAccount || t == entity.BothAccounts {
		if res, err := tx.Exec(ctx, decreaseBalance, updateReq.From(), updateReq.Amount()); err != nil {
			return dao.ResolveError(err)
		} else if res.RowsAffected() == 0 {
			return dao.NewNotFound(updateReq.From())
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return dao.ResolveError(err)
	}

	return nil
}
