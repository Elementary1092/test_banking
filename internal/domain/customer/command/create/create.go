package create

import (
	"context"
	"encoding/hex"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command"
	"github.com/Elementary1092/test_banking/internal/domain/customer/command/model"
	"github.com/Elementary1092/test_banking/pkg/hasher"
	"github.com/google/uuid"
	"time"
)

type Handler struct {
	repo command.WriteDAO
}

func NewCreateCustomerHandler(customerRepo command.WriteDAO) *Handler {
	if customerRepo == nil {
		panic("customer write dao is nil")
	}

	return &Handler{
		repo: customerRepo,
	}
}

func (c *Handler) Handle(ctx context.Context, command Customer) error {
	hashedPassword, err := hasher.Hash(command.Password)
	if err != nil {
		return err
	}
	hexHash := hex.EncodeToString(hashedPassword)

	customer, err := model.NewWriteModel(uuid.New().String(), command.Email, hexHash, time.Now())
	if err != nil {
		return err
	}

	err = c.repo.CreateCustomer(ctx, customer)
	return err
}
