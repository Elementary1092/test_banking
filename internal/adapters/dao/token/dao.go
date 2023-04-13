package token

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	"github.com/Elementary1092/test_banking/internal/entity"
)

type DAO struct {
	db        dao.DB
	tableName string
}

func NewDAO(db dao.DB) *DAO {
	if db == nil {
		panic("db is nil in dao/token.NewDAO")
	}

	return &DAO{
		db:        db,
		tableName: "tokens",
	}
}

func (d *DAO) Insert(ctx context.Context, token *entity.RefreshToken) error {
	const queryFmt = `INSERT INTO %s ("token", "created_at") VALUES ($1, $2)`

	query := fmt.Sprintf(queryFmt, d.tableName)
	if _, err := d.db.Exec(ctx, query, token.Token, token.CreatedAt); err != nil {
		return dao.ResolveError(err)
	}

	return nil
}

func (d *DAO) Delete(ctx context.Context, token string) error {
	const queryFmt = `DELETE FROM %s WHERE "token" = $1`

	query := fmt.Sprintf(queryFmt, d.tableName)
	if _, err := d.db.Exec(ctx, query, token); err != nil {
		return dao.ResolveError(err)
	}

	return nil
}

func (d *DAO) Exists(ctx context.Context, token string) (bool, error) {
	const queryFmt = `SELECT EXISTS(SELECT 1 FROM %s WHERE "token" = $1)`

	query := fmt.Sprintf(queryFmt, d.tableName)
	var exists bool
	if err := d.db.QueryRow(ctx, query, token).Scan(&exists); err != nil {
		return false, dao.ResolveError(err)
	}

	return exists, nil
}
