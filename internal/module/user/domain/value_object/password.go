package value_object

import (
	"strings"

	"github.com/Ali127Dev/xerr"
)

var (
	ErrHashedPasswordRequired = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("hashed_password", "required"))
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
