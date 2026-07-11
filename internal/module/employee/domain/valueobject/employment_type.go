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

func NewEmploymentType(raw string) (EmploymentType, error) {
	switch raw {
	case EmploymentTypeFullTime.value:
		return EmploymentTypeFullTime, nil
	case EmploymentTypePartTime.value:
		return EmploymentTypePartTime, nil
	case EmploymentTypeContract.value:
		return EmploymentTypeContract, nil
	case EmploymentTypeIntern.value:
		return EmploymentTypeIntern, nil
	case EmploymentTypeTemporary.value:
		return EmploymentTypeTemporary, nil
	case EmploymentTypeConsultant.value:
		return EmploymentTypeConsultant, nil
	default:
		return EmploymentType{}, ErrInvalidEmploymentType
	}
}

func (e EmploymentType) Value() string {
	return e.value
}
