package valueobject

import "strings"

type EmployeeID struct {
	value string
}

func NewEmployeeID(id string) EmployeeID {
	id = strings.TrimSpace(id)
	return EmployeeID{value: id}
}

func (u EmployeeID) Value() string { return u.value }
