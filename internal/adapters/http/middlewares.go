package api

import (
	"github.com/Elementary1092/test_banking/internal"
	httpMiddleware "github.com/Elementary1092/test_banking/internal/adapters/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func SetMiddlewares(router http.Handler, config internal.Config) {
	r := chi.NewRouter()

	jwtMiddleware := httpMiddleware.JWTMiddleware{
		Secret: config.TokenGen.Secret,
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(jwtMiddleware.Middleware)
	r.Use(middleware.NoCache)

	r.Mount("/api", router)
}
