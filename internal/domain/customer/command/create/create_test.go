package create

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command/mocks"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command/model"
	"github.com/Elementary1092/test_banking/pkg/hasher"
	"github.com/golang/mock/gomock"
	"testing"
)

var (
	errInvalidEmail    = errors.New("invalid email")
	errInvalidPassword = errors.New("invalid password")
	errInvalidUserID   = errors.New("invalid user id")
)

func TestHandlerHandle(t *testing.T) {
	controller := gomock.NewController(t)
	mockWriteDAO := mocks.NewMockWriteDAO(controller)

	handler := NewCreateCustomerHandler(mockWriteDAO)

	testParams := []struct {
		cmd Customer
		err error
	}{
		{
			cmd: Customer{
				Email:    "some_email",
				Password: "some_password",
			},
			err: nil,
		},
	}

	for _, tt := range testParams {
		mockWriteDAO.
			EXPECT().
			CreateCustomer(gomock.Any(), gomock.Not(gomock.Nil())).
			DoAndReturn(func(ctx context.Context, customer *model.WriteModel) error {
				if customer == nil {
					errMsg := fmt.Sprintf("on %v got unexpected nil", customer)
					return errors.New(errMsg)
				}

				if customer.UserID() == "" {
					return errInvalidUserID
				}

				if customer.Email() != tt.cmd.Email {
					return errInvalidPassword
				}

				hashedPassword, err := hex.DecodeString(customer.Password())
				if err != nil {
					return err
				}

				if err = hasher.Verify(hashedPassword, tt.cmd.Password); err != nil {
					return err
				}

				return nil
			})

		err := handler.Handle(context.Background(), tt.cmd)
		if !errors.Is(err, tt.err) {
			t.Errorf("expected: %v; got: %v", tt.err, err)
		}
	}
}
