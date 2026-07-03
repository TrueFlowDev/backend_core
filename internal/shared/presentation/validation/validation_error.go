package validation

import (
	"github.com/Ali127Dev/xerr"
	"github.com/go-playground/validator/v10"
)

func ToValidationError(errs validator.ValidationErrors) error {
	options := make([]xerr.ErrorOption, 0, len(errs))
	for _, fieldErr := range errs {
		options = append(options, xerr.WithMeta(fieldErr.Field(), ReasonFromTag(fieldErr)))
	}

	return xerr.New(xerr.CodeBadRequest, options...)
}
