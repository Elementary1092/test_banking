package hasher

import (
	"bytes"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
)

const (
	defaultMem = 1 << 16 // 64 KB, but will be set to 64 MB by argon2
	saltLength = 8
	hashLen    = 32
)

var (
	ErrNoPassword   = errors.New("no password was provided")
	ErrDoesNotMatch = errors.New("provided password does not match desired")
)

func hashWithSalt(salt []byte, password string) []byte {
	hash := argon2.IDKey([]byte(password), salt, 1, defaultMem, 1, hashLen)
	return append(salt, hash...)
}

func Hash(password string) ([]byte, error) {
	if password == "" {
		return nil, ErrNoPassword
	}

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return hashWithSalt(salt, password), nil
}

func Verify(hash []byte, password string) error {
	if password == "" {
		return ErrNoPassword
	}

	// limiting capacity. Otherwise, original hash may be rewritten by a new hash value
	salt := hash[:saltLength:saltLength]
	hashedPassword := hashWithSalt(salt, password)
	if !bytes.Equal(hashedPassword, hash) {
		return ErrDoesNotMatch
	}

	return nil
}
