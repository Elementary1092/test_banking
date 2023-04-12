package hasher

import (
    "crypto/rand"
    "encoding/hex"
    "errors"

    "golang.org/x/crypto/argon2"
)

const (
    defaultMem = 1 << 16 // 64 KB, but will be set to 64 MB by argon2
    saltLength = 8
    hashLen    = 32
)

var (
    ErrNoPassword       = errors.New("no password was provided")
    ErrNoHashedPassword = errors.New("no hashed password was provided")
    ErrInvalidHash      = errors.New("provided hashed password is invalid")
    ErrDoesNotMatch     = errors.New("provided password does not match desired")
)

func hashWithSalt(salt []byte, password string) string {
    hash := argon2.IDKey([]byte(password), salt, 1, defaultMem, 1, hashLen)
    hashedPass := append(salt, hash...)

    return hex.EncodeToString(hashedPass)
}

func Hash(password string) (string, error) {
    if password == "" {
        return "", ErrNoPassword
    }

    salt := make([]byte, saltLength)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }

    return hashWithSalt(salt, password), nil
}

func Verify(hashedPassword, password string) error {
    if password == "" {
        return ErrNoPassword
    }
    if hashedPassword == "" {
        return ErrNoHashedPassword
    }

    hash, err := hex.DecodeString(hashedPassword)
    if err != nil {
        return ErrInvalidHash
    }

    // limiting capacity. Otherwise, original hash may be rewritten by a new hash value
    salt := hash[:saltLength:saltLength]
    testingPassword := hashWithSalt(salt, password)
    if testingPassword != hashedPassword {
        return ErrDoesNotMatch
    }

    return nil
}
