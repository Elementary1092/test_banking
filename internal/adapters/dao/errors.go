package dao

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

var (
	ErrDatabaseOverload = errors.New("database is overloaded")
	ErrInternalError    = errors.New("internal error")
	ErrMemoryLimit      = errors.New("memory limit exceeded")
	ErrComplexStmt      = errors.New("statement is too complex")
	ErrInvalidStmt      = errors.New("invalid statement")
)

type ErrInvalidEmpty struct {
	columnName string
}

func (e ErrInvalidEmpty) Error() string {
	return fmt.Sprintf("%s cannot be empty", e.columnName)
}

type ErrNotFound struct {
	data string
}

func NewNotFound(data string) ErrNotFound {
	return ErrNotFound{data: data}
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s does not exist", e.data)
}

type ErrDuplicate struct {
	data string
}

func (e ErrDuplicate) Error() string {
	return fmt.Sprintf("%s already exists", e.data)
}

func ResolveError(err error) error {
	var pgErr *pgconn.PgError
	if !errors.Is(err, pgErr) {
		return ErrInternalError
	}

	pgErr = err.(*pgconn.PgError)
	switch pgErr.Code {
	case "53200": // out of memory
		return ErrMemoryLimit

	case "53300":
		return ErrDatabaseOverload

	case "54001": // statement is too complex
		return ErrComplexStmt

	// 54011 - too many columns
	// 54023 - too many arguments
	// 42P10 - invalid column reference
	// 42703 - undefined column
	// 42P01 - undefined table
	case "54011", "54023", "42P10", "42703", "42P01":
		return ErrInvalidStmt

	case "23502": // NOT NULL violation
		return ErrInvalidEmpty{columnName: pgErr.ColumnName}

	case "23503", "P0002": // FOREIGN KEY violation, NO DATA FOUND
		// TODO add data parsing from pgErr.Details
		return ErrNotFound{data: "some data"}

	case "23505":
		// TODO add data parsing from pgErr.Details
		return ErrDuplicate{data: "data"}

	default:
		return ErrInternalError
	}
}
