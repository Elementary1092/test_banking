package create

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	accountNumberBufLen             = 10
	expectedNumericAccountNumberLen = 20
)

var (
	errFailedToGenerateSufficientNumber = errors.New("internal error occurred during account number generation")
	errInvalidCurrencyValue             = errors.New("invalid currency value was provided")
)

func generateAccountNumber(currency string) (string, error) {
	if currency == "" {
		return "", errInvalidCurrencyValue
	}

	// Yeah, it is not the best solution to create account number using crypto/rand.Read
	// but, for now, I think it is ok.
	//
	// 10 for buffer is size is chosen to get a string with 20 characters after hex.EncodeToString() call.
	accountNumberBuf := make([]byte, accountNumberBufLen)

	if n, err := rand.Read(accountNumberBuf); err != nil {
		return "", err
	} else if n != accountNumberBufLen {
		return "", errFailedToGenerateSufficientNumber
	}

	accountNumber := hex.EncodeToString(accountNumberBuf)

	return fmt.Sprintf("%s-%s", currency, accountNumber), nil
}
