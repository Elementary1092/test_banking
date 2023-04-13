package handler

import (
	"github.com/Elementary1092/test_banking/internal/adapters/http/httperr"
	"net/http"
)

func Render(w http.ResponseWriter, data any, successStat int) {
	encoder := json.NewEncoder(w)

	if err := encoder.Encode(data); err != nil {
		httperr.InternalError(w, "failed-to-write-response")
		return
	}

	w.WriteHeader(successStat)
}
