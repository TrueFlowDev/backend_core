package validation

import (
	"github.com/Ali127Dev/xerr"
	"github.com/go-playground/validator/v10"
)

func ReasonFromTag(err validator.FieldError) xerr.ErrorReason {
	switch err.Tag() {
	case "required":
		return xerr.ErrorReasonRequired

	case "min":
		return xerr.ErrorReasonTooShort

	case "max":
		return xerr.ErrorReasonTooLong

	case "email", "e164", "uuid", "url", "oneof":
		return xerr.ErrorReasonInvalidFormat

	case "eqfield", "nefield":
		return xerr.ErrorReasonMismatch

	default:
		return xerr.ErrorReasonInvalidFormat
	}
}
