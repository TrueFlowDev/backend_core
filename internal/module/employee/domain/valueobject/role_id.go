package valueobject

import "strings"

type RoleID struct {
	value string
}

func NewRoleID(id string) RoleID {
	id = strings.TrimSpace(id)
	return RoleID{value: id}
}

func (u RoleID) Value() string { return u.value }
