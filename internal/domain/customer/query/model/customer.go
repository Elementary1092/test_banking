package model

import (
	"github.com/Elementary1092/test_banking/internal/entity"
	"time"
)

type Customer struct {
	user *entity.User

	// will remain read-only
	createdAt time.Time
}

func NewCustomer(userID, email, password string, createdAt time.Time) *Customer {
	return &Customer{
		user: &entity.User{
			UUID:     userID,
			Email:    email,
			Password: password,
		},
		createdAt: createdAt,
	}
}

func (c *Customer) UUID() string {
	return c.user.UUID
}

func (c *Customer) SetUUID(uuid string) {
	c.user.UUID = uuid
}

func (c *Customer) Email() string {
	return c.user.Email
}

func (c *Customer) SetEmail(email string) {
	c.user.Email = email
}

func (c *Customer) Password() string {
	return c.user.Password
}

func (c *Customer) SetPassword(password string) {
	c.user.Password = password
}

func (c *Customer) CreatedAt() time.Time {
	return c.createdAt
}
