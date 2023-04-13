package server

import (
	"net/http"
	"strings"
)

func New(config Config) http.Server {
	return http.Server{
		Addr:         parseAddress(config.Address),
		Handler:      config.Router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
}

func parseAddress(address string) string {
	address = strings.TrimSpace(address)

	// This piece of code handles address where port without ':' is handed over
	if !strings.Contains(address, ":") {
		return ":" + address
	}

	return address
}
