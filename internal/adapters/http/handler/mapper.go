package handler

import (
	"errors"
	api "github.com/Elementary1092/test_banking/internal/adapters/http"
	accountRead "github.com/Elementary1092/test_banking/internal/domain/account/query/model"
	customerRead "github.com/Elementary1092/test_banking/internal/domain/customer/query/model"
	"github.com/google/uuid"
)

var (
	ErrInvalidCurrency = errors.New("invalid currency")
)

func MapCustomerRead(customer *customerRead.Customer) api.Customer {
	createdAt := customer.CreatedAt()
	email := customer.Email()
	uuidStr := customer.UUID()
	userID := uuid.MustParse(uuidStr)

	return api.Customer{
		CreatedAt: &createdAt,
		Email:     &email,
		Uuid:      &userID,
	}
}

func MapSignInResponse(idToken, refreshToken string) api.SignInResponse {
	return api.SignInResponse{
		IdToken:      &idToken,
		RefreshToken: &refreshToken,
	}
}

func MapAccountRead(account *accountRead.Account) api.GetAccountResponse {
	userID := uuid.MustParse(account.UserID)

	return api.GetAccountResponse{
		AccountNumber: &account.Info.Number,
		Balance:       &account.Info.Balance,
		Currency:      &account.Info.Currency,
		CustomerId:    &userID,
	}
}
