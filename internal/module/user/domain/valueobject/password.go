package valueobject

import (
	"strings"

	"github.com/Ali127Dev/xerr"
)

var (
	ErrHashedPasswordRequired = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("password", xerr.ErrorReasonRequired))
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
