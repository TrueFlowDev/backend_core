package value_object

import (
	"errors"
	"strings"
)

var (
	ErrHashedPasswordRequired = errors.New("hashed password is required")
)

type HashedPassword struct {
	value string
}

func NewHashedPassword(hash string) (HashedPassword, error) {
	hash = strings.TrimSpace(hash)

	if hash == "" {
		return HashedPassword{}, ErrHashedPasswordRequired
	}

	return HashedPassword{value: hash}, nil
}

func (p HashedPassword) Value() string { return p.value }
