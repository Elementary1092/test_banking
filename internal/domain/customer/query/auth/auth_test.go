package auth

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/mocks"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"github.com/Elementary1092/test_banking/pkg/hasher"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

var (
	errInvalidParams = errors.New("invalid search parameters were received")
	errNotFound      = errors.New("customer was not found")
)

func TestAuthCustomer(t *testing.T) {
	controller := gomock.NewController(t)
	mockReadDAO := mocks.NewMockReadDAO(controller)

	finder := NewAuthHandler(mockReadDAO)
	tests := map[Query]struct {
		err      error
		customer *model.Customer
	}{}

	expectedPassword := "some_password"
	expectedPasswordHash, err := hasher.Hash(expectedPassword)
	if err != nil {
		t.Errorf("failed on generating hash of an expected password: %v", err)
	}

	unexpectedPassword := "other_password"

	createTime := time.Now().Truncate(10 * time.Minute)
	expectedCustomer := model.NewCustomer("user_id", "some_email", expectedPasswordHash, createTime)

	queryWithValidData := Query{
		Email:    expectedCustomer.Email(),
		Password: expectedPassword,
	}
	customerWithValidDataResponse := model.NewCustomer("user_id", "some_email", expectedPasswordHash, createTime)
	tests[queryWithValidData] = struct {
		err      error
		customer *model.Customer
	}{
		err:      nil,
		customer: customerWithValidDataResponse,
	}

	queryWithInvalidEmail := Query{
		Email:    "some_invalid_email",
		Password: expectedPassword,
	}
	tests[queryWithInvalidEmail] = struct {
		err      error
		customer *model.Customer
	}{
		err:      ErrInvalidAuthParameters,
		customer: nil,
	}

	queryWithInvalidPassword := Query{
		Email:    expectedCustomer.Email(),
		Password: unexpectedPassword,
	}
	tests[queryWithInvalidPassword] = struct {
		err      error
		customer *model.Customer
	}{
		err:      ErrInvalidAuthParameters,
		customer: nil,
	}

	mockReadDAO.
		EXPECT().
		FindCustomer(gomock.Any(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(ctx context.Context, params map[string]string) (*model.Customer, error) {
			if len(params) == 0 {
				return nil, errInvalidParams
			}

			for key, value := range params {
				switch key {
				case "uuid":
					if value == expectedCustomer.UUID() {
						return expectedCustomer, nil
					}

				case "email":
					if value == expectedCustomer.Email() {
						return expectedCustomer, nil
					}
				}
			}

			return nil, errNotFound
		}).
		Times(3) // will be called <number positive results> times

	for query, result := range tests {
		customer, err := finder.Handle(context.Background(), query)
		if !errors.Is(err, result.err) {
			t.Errorf("exptected error: %v; got: %v", result.err, err)
		}

		if !reflect.DeepEqual(customer, result.customer) {
			t.Errorf("unexpected result on %v", *(result.customer))
		}
	}
}
