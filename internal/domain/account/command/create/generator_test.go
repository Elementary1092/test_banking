package create

import (
	"errors"
	"testing"
)

func TestAccountNumberGenerator(t *testing.T) {
	tests := map[string]error{
		"UZS": nil,
		"USD": nil,
		"":    errInvalidCurrencyValue,
	}

	for currency, expectedErr := range tests {
		generated, err := generateAccountNumber(currency)
		if !errors.Is(err, expectedErr) {
			t.Errorf("expected error: %v; got: %v | on %s", expectedErr, err, currency)
		}

		if err == nil && len(generated) != len(currency)+expectedNumericAccountNumberLen+1 {
			t.Errorf(
				"got account number with invalid length: expected %d, got %d",
				len(currency)+2*accountNumberBufLen+1,
				len(generated),
			)
		}
	}
}
