package phonenumber

import (
	"strings"

	"github.com/Ali127Dev/xerr"
	"github.com/nyaruka/phonenumbers"
)

var (
	ErrRequired      = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("phone", "required"))
	ErrInvalidFormat = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("phone", "invalid"))
)

func NormalizePhone(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", ErrRequired
	}

	number, err := phonenumbers.Parse(raw, "IR")
	if err != nil {
		return "", ErrInvalidFormat
	}

	if !phonenumbers.IsValidNumber(number) {
		return "", ErrInvalidFormat
	}

	if phonenumbers.GetNumberType(number) != phonenumbers.MOBILE {
		return "", ErrInvalidFormat
	}

	return phonenumbers.Format(number, phonenumbers.E164), nil
}
