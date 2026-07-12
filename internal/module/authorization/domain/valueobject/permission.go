package valueobject

import (
	"sort"

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
	RolePermissions,
)

type PermissionCategory string

type Permission struct {
	category PermissionCategory
	title    string
	value    string
}

func ParsePermission(raw string) (Permission, error) {
	p, ok := allPermissions[raw]
	if !ok {
		return Permission{}, ErrInvalidPermission
	}

	return p, nil
}

func (p Permission) Category() PermissionCategory { return p.category }
func (p Permission) Title() string                { return p.title }
func (p Permission) Value() string                { return p.value }
func All() []Permission {
	result := make([]Permission, 0, len(allPermissions))
	for _, p := range allPermissions {
		result = append(result, p)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].value < result[j].value
	})
	return result
}

func (c PermissionCategory) String() string { return string(c) }

func mergePermissions(maps ...map[string]Permission) map[string]Permission {
	result := make(map[string]Permission)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func buildPermissionMap(perms ...Permission) map[string]Permission {
	result := make(map[string]Permission, len(perms))
	for _, p := range perms {
		result[p.value] = p
	}
	return result
}
