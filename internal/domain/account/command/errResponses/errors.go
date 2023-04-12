package errResponses

import (
	"errors"
	"fmt"
)

var (
	ErrFailedToGenAccNumber = errors.New("account number generation error")
	ErrInvalidCurrency      = errors.New("invalid currency value")

	ErrInvalidAccountInfo = errors.New("invalid account information")
	ErrInvalidBalanceInfo = errors.New("invalid account balance information")

	ErrInvalidUpdateParameters  = errors.New("invalid update parameters")
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
	ErrInsufficientFunds        = errors.New("insufficient balance")

	ErrInvalidRecipientCurrency = errors.New("recipient's account currency is different")
)

type errInternalError struct {
	failMsg string
	wrapped error
}

func (e *errInternalError) Error() string {
	return fmt.Sprintf("Failed on %s. Actual error: %v", e.failMsg, e.wrapped)
}

func (e *errInternalError) Unwrap() error {
	return e.wrapped
}

func NewInternal(opDesc string, err error) error {
	return &errInternalError{
		failMsg: opDesc,
		wrapped: err,
	}
}
