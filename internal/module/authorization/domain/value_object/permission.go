package value_object

import (
	"github.com/Ali127Dev/xerr"
)

var (
	ErrInvalidPermission = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("permission", xerr.ErrorReasonInvalidFormat),
	)
)

var allPermissions = mergePermissions(
	OrganizationPermissions,
	EmployeePermissions,
)

type Permission struct {
	title string
	value string
}

func NewPermission(raw string) (Permission, error) {
	p, ok := allPermissions[raw]
	if !ok {
		return Permission{}, ErrInvalidPermission
	}

	return p, nil
}

func (p Permission) Title() string { return p.title }
func (p Permission) Value() string { return p.value }

func mergePermissions(maps ...map[string]Permission) map[string]Permission {
	result := make(map[string]Permission)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}
