package valueobject

import (
	"github.com/Ali127Dev/xerr"
)

var (
	ErrInvalidEmploymentType = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("employment_type", xerr.ErrorReasonInvalidFormat),
	)
)

type EmploymentType struct {
	value string
}

var (
	EmploymentTypeFullTime   = EmploymentType{"full_time"}
	EmploymentTypePartTime   = EmploymentType{"part_time"}
	EmploymentTypeContract   = EmploymentType{"contract"}
	EmploymentTypeIntern     = EmploymentType{"intern"}
	EmploymentTypeTemporary  = EmploymentType{"temporary"}
	EmploymentTypeConsultant = EmploymentType{"consultant"}
)

var employmentTypes = buildEmploymentTypeMap(
	EmploymentTypeFullTime,
	EmploymentTypePartTime,
	EmploymentTypeContract,
	EmploymentTypeIntern,
	EmploymentTypeTemporary,
	EmploymentTypeConsultant,
)

func ParseEmploymentType(raw string) (EmploymentType, error) {
	e, ok := employmentTypes[raw]
	if !ok {
		return EmploymentType{}, ErrInvalidEmploymentType
	}
	return e, nil
}

func (e EmploymentType) Value() string {
	return e.value
}

func buildEmploymentTypeMap(types ...EmploymentType) map[string]EmploymentType {
	m := make(map[string]EmploymentType, len(types))
	for _, t := range types {
		m[t.value] = t
	}
	return m
}
