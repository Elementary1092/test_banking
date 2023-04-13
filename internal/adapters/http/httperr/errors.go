package httperr

import (
	"errors"
	"github.com/Elementary1092/test_banking/internal/adapters/dao"
	accCmdErr "github.com/Elementary1092/test_banking/internal/domain/account/command/errResponses"
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ErrorResponse struct {
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}

// WrapError wraps errors from command, queries and DAO
func WrapError(w http.ResponseWriter, err error) {
	resp := ErrorResponse{
		Status: http.StatusBadRequest,
	}

	// Not the best decision, unique error codes could have better
	if errors.Is(err, accCmdErr.ErrInsufficientFunds) {
		resp.Slug = "insufficient-balance"
		resp.Status = http.StatusBadRequest
	} else if errors.Is(err, accCmdErr.ErrInvalidTransactionAmount) {
		resp.Slug = "invalid-amount"
	} else if errors.Is(err, accCmdErr.ErrInvalidAccountInfo) {
		resp.Slug = "invalid-account-info"
	} else if errors.Is(err, accCmdErr.ErrFailedToGenAccNumber) {
		resp.Status = http.StatusInternalServerError
		resp.Slug = "account-number-error"
	} else if errors.Is(err, accCmdErr.ErrInvalidRecipientCurrency) {
		resp.Slug = "currencies-do-not-match"
	} else if errors.Is(err, accCmdErr.ErrInvalidUpdateParameters) {
		resp.Slug = "invalid-account-data"
	} else if errors.Is(err, accCmdErr.ErrInvalidBalanceInfo) {
		resp.Slug = "invalid-balance"
	} else if errors.Is(err, dao.ErrDuplicate{}) {
		resp.Slug = "duplicate-data"
		resp.Status = http.StatusConflict
	} else if errors.Is(err, dao.ErrInvalidEmpty{}) {
		resp.Slug = "missing-required-data"
	} else if errors.Is(err, dao.ErrNotFound{}) {
		resp.Slug = "not-found"
		resp.Status = http.StatusNotFound
	} else {
		resp.Slug = "internal-error"
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(resp.Status)
	encoder.Encode(resp)
}

func Unauthorized(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusUnauthorized,
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusUnauthorized)
	encoder.Encode(resp)
}

func InternalError(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusInternalServerError,
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusInternalServerError)
	encoder.Encode(resp)
}

func BadRequest(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusBadRequest,
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusBadRequest)
	encoder.Encode(resp)
}
