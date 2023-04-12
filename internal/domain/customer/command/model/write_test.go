package model

import (
	"errors"
	"testing"
	"time"
)

func TestNewWriteModel(t *testing.T) {
	type writeModelParams struct {
		email, password, uuid string
		createdAt             time.Time
	}

	tests := map[writeModelParams]error{
		{
			uuid:      "unique_id",
			email:     "some_email",
			password:  "some_password",
			createdAt: time.Now(),
		}: nil,
		{
			email:     "some_email",
			password:  "some_password",
			createdAt: time.Now(),
		}: ErrInvalidCustomer,
		{
			uuid:      "unique_id",
			email:     "some_email",
			createdAt: time.Now(),
		}: ErrInvalidCustomer,
		{
			uuid:      "unique_id",
			email:     "some_email",
			password:  "some_password",
			createdAt: time.Now().Add(5 * time.Hour),
		}: ErrInvalidCustomer,
	}

	for params, expectedErr := range tests {
		_, err := NewWriteModel(params.uuid, params.email, params.password, params.createdAt)
		if !errors.Is(err, expectedErr) {
			t.Errorf("expected error: %v: got: %v\nValues: %v", expectedErr, err, params)
		}
	}
}
