package token

import (
	"context"
	"errors"
	"github.com/Elementary1092/test_banking/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	refreshTokenAddOn = "refresh_token"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
	ErrExpired              = errors.New("token is expired")
)

type Handler struct {
	repo              DAO
	refreshExpiration time.Duration
	idExpiration      time.Duration
	issuer            string
	secret            string
}

func NewHandler(
	repo DAO,
	refreshExpiration, idExpiration time.Duration,
	issuer, secret string) *Handler {
	if repo == nil {
		panic("repo is nil in token.NewHandler()")
	}

	return &Handler{
		repo:              repo,
		refreshExpiration: refreshExpiration,
		idExpiration:      idExpiration,
		issuer:            issuer,
		secret:            secret,
	}
}

func (h *Handler) GenerateToken(ctx context.Context, userID string) (string, string, error) {
	issueTime := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    h.issuer,
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(issueTime.Add(h.idExpiration)),
		IssuedAt:  jwt.NewNumericDate(issueTime),
	}

	idToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(h.secret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.RegisteredClaims{
		Issuer:    h.issuer,
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(issueTime.Add(h.refreshExpiration)),
		IssuedAt:  jwt.NewNumericDate(issueTime),
		ID:        refreshTokenAddOn,
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(h.secret)
	if err != nil {
		return "", "", err
	}

	err = h.repo.Insert(ctx, &entity.RefreshToken{
		Token:     refreshToken,
		CreatedAt: issueTime,
	})
	if err != nil {
		return "", "", err
	}

	return idToken, refreshToken, nil
}

func (h *Handler) RefreshToken(ctx context.Context, tokenStr string) (string, string, error) {
	if ok, err := h.repo.Exists(ctx, tokenStr); err != nil || !ok {
		return "", "", ErrInvalidToken
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrInvalidSigningMethod
		}

		return []byte(h.secret), nil
	})
	if err != nil {
		if errors.Is(err, ErrInvalidSigningMethod) {
			return "", "", err
		}
		return "", "", ErrInvalidToken
	}

	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return "", "", ErrInvalidToken
	}

	if exp.Time.Before(time.Now()) {
		return "", "", ErrExpired
	}

	if err = h.repo.Delete(ctx, tokenStr); err != nil {
		return "", "", err
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return "", "", ErrInvalidToken
	}

	return h.GenerateToken(ctx, sub)
}
