package valueobject

import (
	"net/mail"
	"strings"

	"github.com/Ali127Dev/xerr"
)

var (
	ErrEmailRequired      = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("email", xerr.ErrorReasonRequired))
	ErrEmailInvalidFormat = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("email", xerr.ErrorReasonInvalidFormat))
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return Email{}, ErrEmailRequired
	}

	email = strings.ToLower(email)

	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return Email{}, ErrEmailInvalidFormat
	}

	return Email{value: email}, nil
}

func (e Email) Value() string { return e.value }
