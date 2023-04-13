package middleware

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/adapters/http/httperr"
	domainToken "github.com/Elementary1092/test_banking/internal/domain/token"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var (
	ErrNoTokenInfo = errors.New("failed to get token info")
)

type JWTMiddleware struct {
	Secret string
}

type ctxKey int

const (
	userCtxKey ctxKey = iota
)

func (m JWTMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenStr := tokenFromHeader(r)
		if tokenStr == "" {
			httperr.Unauthorized(w, "no-token")
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", domainToken.ErrInvalidSigningMethod
			}

			return []byte(m.Secret), nil
		})
		if err != nil {
			httperr.Unauthorized(w, "invalid-token")
			return
		}

		expTimeStr, err := token.Claims.GetExpirationTime()
		if err != nil {
			httperr.Unauthorized(w, "no-expr-time")
			return
		}
		if expTimeStr.Time.Before(time.Now()) {
			httperr.Unauthorized(w, "token-expired")
			return
		}

		userID, err := token.Claims.GetSubject()
		if err != nil {
			httperr.Unauthorized(w, "no-token-subject")
			return
		}
		ctx = context.WithValue(ctx, userCtxKey, map[string]string{
			"user_id": userID,
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func RetrieveUserID(ctx context.Context) (string, error) {
	userInfo, ok := ctx.Value(userCtxKey).(map[string]string)
	if !ok {
		return "", ErrNoTokenInfo
	}

	val, ok := userInfo["user_id"]
	if !ok {
		return "", ErrNoTokenInfo
	}

	return val, nil
}

func tokenFromHeader(r *http.Request) string {
	fromHeader := r.Header.Get("Authorization")

	if len(fromHeader) < 7 || strings.ToLower(fromHeader[:6]) != "bearer" {
		return ""
	}

	return fromHeader[7:]
}
