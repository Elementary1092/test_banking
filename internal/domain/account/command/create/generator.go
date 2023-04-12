package create

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
)

const (
	accountNumberBufLen             = 10
	expectedNumericAccountNumberLen = 20
)

func generateAccountNumber(currency string) (string, error) {
	if currency == "" {
		return "", errResponses.ErrInvalidCurrency
	}

	// Yeah, it is not the best solution to create account number using crypto/rand.Read
	// but, for now, I think it is ok.
	//
	// 10 for buffer is size is chosen to get a string with 20 characters after hex.EncodeToString() call.
	accountNumberBuf := make([]byte, accountNumberBufLen)

	if n, err := rand.Read(accountNumberBuf); err != nil {
		return "", errResponses.NewInternal("crypto/rand.Read", err)
	} else if n != accountNumberBufLen {
		return "", errResponses.ErrFailedToGenAccNumber
	}

	accountNumber := hex.EncodeToString(accountNumberBuf)

	return fmt.Sprintf("%s-%s", currency, accountNumber), nil
}
