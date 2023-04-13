package httperr

import (
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ErrorResponse struct {
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}

func WrapError(w http.ResponseWriter, err error) {
	resp := ErrorResponse{
		Slug:   err.Error(),
		Status: http.StatusBadRequest,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
	w.WriteHeader(http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusUnauthorized,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
	w.WriteHeader(http.StatusUnauthorized)
}

func InternalError(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusInternalServerError,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
	w.WriteHeader(http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, slug string) {
	resp := ErrorResponse{
		Slug:   slug,
		Status: http.StatusBadRequest,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
	w.WriteHeader(http.StatusBadRequest)
}
