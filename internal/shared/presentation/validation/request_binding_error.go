package validation

import (
	"fmt"

	"github.com/Ali127Dev/xerr"
)

type Kind string

const (
	JSON  Kind = "json"
	YAML  Kind = "yaml"
	FORM  Kind = "form"
	QUERY Kind = "query"
	PARAM Kind = "param"
)

func NewRequestBindingError(handler string, kind Kind) error {
	return xerr.Wrap(fmt.Errorf("failed to bind %s %s", handler, kind), xerr.CodeBadRequest)
}
