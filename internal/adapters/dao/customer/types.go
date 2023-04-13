package customer

import (
	"github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"time"
)

type customerModel struct {
	UUID      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (c customerModel) MapToReadModel() *model.Customer {
	return model.NewCustomer(c.UUID, c.Email, c.Password, c.CreatedAt)
}
