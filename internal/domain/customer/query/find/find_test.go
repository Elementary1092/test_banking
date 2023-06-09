package find

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/mocks"
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

var (
	errInvalidParams = errors.New("invalid search parameters were received")
	errNotFound      = errors.New("customer was not found")
)

func TestFindCustomer(t *testing.T) {
	controller := gomock.NewController(t)
	mockReadDAO := mocks.NewMockReadDAO(controller)

	finder := NewHandler(mockReadDAO)
	tests := map[Query]struct {
		err      error
		customer *model.Customer
	}{}

	createTime := time.Now().Truncate(10 * time.Minute)
	expectedCustomer := model.NewCustomer("user_id", "some_email", "", createTime)

	customerWithEmail := Query{
		Email: "some_email",
	}
	customerWithEmailResponse := model.NewCustomer("user_id", "some_email", "", createTime)
	tests[customerWithEmail] = struct {
		err      error
		customer *model.Customer
	}{
		err:      nil,
		customer: customerWithEmailResponse,
	}

	customerWithUUID := Query{
		UUID: "user_id",
	}
	customerWithUUIDResponse := model.NewCustomer("user_id", "some_email", "", createTime)
	tests[customerWithUUID] = struct {
		err      error
		customer *model.Customer
	}{
		err:      nil,
		customer: customerWithUUIDResponse,
	}

	mockReadDAO.
		EXPECT().
		FindCustomer(gomock.Any(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(ctx context.Context, params map[string]string) (*model.Customer, error) {
			if len(params) == 0 {
				return nil, errInvalidParams
			}

			var resultCustomer *model.Customer
			var err error
			for key, value := range params {
				switch key {
				case "uuid":
					if value == expectedCustomer.UUID() {
						resultCustomer = expectedCustomer
					} else if resultCustomer != nil && err == nil {
						resultCustomer = nil
						err = errNotFound
					}

				case "email":
					if value == expectedCustomer.Email() {
						resultCustomer = expectedCustomer
					} else if resultCustomer != nil && err == nil {
						resultCustomer = nil
						err = errNotFound
					}
				}
			}

			return resultCustomer, err
		}).
		Times(2) // will be called <number positive results> times

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
