package hasher

import (
	"errors"
	"fmt"
	"testing"
)

func defaultResultChecker(password string) func([]byte, error) error {
	return func(result []byte, err error) error {
		if err != nil {
			errMsg := fmt.Sprintf("unexpected error on '%s': %v", password, err)
			return errors.New(errMsg)
		}
		if result == nil || len(result) < saltLength+hashLen {
			errMsg := fmt.Sprintf("invalid hash on '%s'", password)
			return errors.New(errMsg)
		}

		return nil
	}
}

func defaultErrorChecker(password string, desiredError error) func([]byte, error) error {
	return func(result []byte, err error) error {
		if !errors.Is(err, desiredError) {
			errMsg := fmt.Sprintf("failed on: %s; expected: %v, got: %v", password, desiredError, err)
			return errors.New(errMsg)
		}

		return nil
	}
}

func TestHash(t *testing.T) {
	tests := map[string]func(result []byte, err error) error{
		"some_password":             defaultResultChecker("some_password"),
		"":                          defaultErrorChecker("", ErrNoPassword),
		"1":                         defaultResultChecker("1"),
		"some_really_long_password": defaultResultChecker("some_really_long_password"),
	}

	for password, checker := range tests {
		result, errHash := Hash(password)
		if err := checker(result, errHash); err != nil {
			t.Error(err)
		}
	}
}

func TestVerify(t *testing.T) {
	tests := map[string]struct {
		compareWith       string
		expectedVerifyErr error
		expectedHashErr   error
	}{
		"some_password":      {compareWith: "some_password"},
		"":                   {compareWith: "", expectedVerifyErr: ErrNoPassword, expectedHashErr: ErrNoPassword},
		"1":                  {compareWith: "2", expectedVerifyErr: ErrDoesNotMatch},
		"some_long_password": {compareWith: "other_long_password", expectedVerifyErr: ErrDoesNotMatch},
	}

	for password, desc := range tests {
		hash, err := Hash(password)
		errHashChecker := defaultErrorChecker(password, desc.expectedHashErr)
		if err = errHashChecker(nil, err); err != nil {
			t.Error(err)
		}

		err = Verify(hash, desc.compareWith)
		errVerifyChecker := defaultErrorChecker(password, desc.expectedVerifyErr)
		if err = errVerifyChecker(nil, err); err != nil {
			t.Error(err)
		}
	}
}
