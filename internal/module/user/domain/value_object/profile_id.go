package value_object

import "strings"

type ProfileID struct {
	value string
}

func NewProfileID(id string) ProfileID {
	id = strings.TrimSpace(id)
	return ProfileID{value: id}
}

func (u ProfileID) Value() string { return u.value }
