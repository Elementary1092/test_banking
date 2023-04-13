package entity

import "time"

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
}
