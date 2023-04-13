package model

import (
	"errors"
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

var (
	ErrInvalidCustomer = errors.New("invalid parameters set to create customer")
)

type WriteModel struct {
	createdAt time.Time
	user      *entity.User
}

func NewWriteModel(customerID, email, password string, createdAt time.Time) (*WriteModel, error) {
	if customerID == "" || email == "" || password == "" {
		return nil, ErrInvalidCustomer
	}

	if createdAt.After(time.Now()) {
		return nil, ErrInvalidCustomer
	}

	return &WriteModel{
		user: &entity.User{
			UUID:     customerID,
			Email:    email,
			Password: password,
		},
		createdAt: createdAt,
	}, nil
}

func (c *WriteModel) UserID() string {
	return c.user.UUID
}

func (c *WriteModel) Email() string {
	return c.user.Email
}

func (c *WriteModel) Password() string {
	return c.user.Password
}

func (c *WriteModel) CreatedAt() time.Time {
	return c.createdAt
}
