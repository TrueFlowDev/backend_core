package valueobject

import "strings"

type OrganizationID struct {
	value string
}

func NewOrganizationID(id string) OrganizationID {
	id = strings.TrimSpace(id)
	return OrganizationID{value: id}
}

func (u OrganizationID) Value() string { return u.value }
