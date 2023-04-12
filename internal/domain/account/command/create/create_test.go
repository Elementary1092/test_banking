package create

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/mocks"
	"github.com/Elementary1092/test_banking/internal/domain/account/command/model"
	"github.com/golang/mock/gomock"
	"testing"
)

var (
	errInvalidAccount = errors.New("invalid account")
)

func TestHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockWriteDAO := mocks.NewMockWriteDAO(controller)

	handler := NewCreateHandler(mockWriteDAO)

	mockWriteDAO.
		EXPECT().
		CreateAccount(gomock.Any(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(ctx context.Context, account *model.Account) error {
			if account == nil ||
				account.UserID() == "" ||
				account.Balance() != 0 ||
				account.Currency() == "" ||
				len(account.Number()) < expectedNumericAccountNumberLen {

				return errInvalidAccount
			}
			return nil
		}).
		Times(2)

	tests := []Command{
		{
			UserID:   "some_user_id",
			Currency: "UZS",
		},
		{
			UserID:   "other_user_id",
			Currency: "USD",
		},
	}

	for _, cmd := range tests {
		err := handler.Handle(context.Background(), cmd)
		if err != nil {
			t.Errorf("unexpected error on %v: %v", cmd, err)
		}
	}
}
