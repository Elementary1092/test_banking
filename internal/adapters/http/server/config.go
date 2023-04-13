package server

import (
	"net/http"
	"time"
)

type Config struct {
	Address      string
	Router       http.Handler
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
