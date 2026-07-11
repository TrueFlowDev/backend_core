package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/value_object"
)

const (
	MinJobTitleLength = 2
	MaxJobTitleLength = 100
)

var (
	ErrJobTitleTooShort = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("job_title", xerr.ErrorReasonTooShort),
	)

	ErrJobTitleTooLong = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("job_title", xerr.ErrorReasonTooLong),
	)
)

type Employee struct {
	id             value_object.EmployeeID
	userID         value_object.UserID
	organizationID value_object.OrganizationID

	jobTitle         string
	membershipStatus value_object.MembershipStatus
	employmentType   value_object.EmploymentType

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreEmployeeParams struct {
	ID             value_object.EmployeeID
	UserID         value_object.UserID
	OrganizationID value_object.OrganizationID

	JobTitle         string
	MembershipStatus value_object.MembershipStatus
	EmploymentType   value_object.EmploymentType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewEmployee(
	id value_object.EmployeeID,
	userID value_object.UserID,
	organizationID value_object.OrganizationID,
	jobTitle string,
	membershipStatus value_object.MembershipStatus,
	employmentType value_object.EmploymentType,
) (*Employee, error) {
	jobTitle, err := validateJobTitle(jobTitle)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Employee{
		id:               id,
		userID:           userID,
		organizationID:   organizationID,
		jobTitle:         jobTitle,
		membershipStatus: membershipStatus,
		employmentType:   employmentType,
		createdAt:        now,
		updatedAt:        now,
	}, nil
}

func RestoreEmployee(
	params RestoreEmployeeParams,
) *Employee {
	return &Employee{
		id:               params.ID,
		userID:           params.UserID,
		organizationID:   params.OrganizationID,
		jobTitle:         params.JobTitle,
		membershipStatus: params.MembershipStatus,
		employmentType:   params.EmploymentType,
		createdAt:        params.CreatedAt,
		updatedAt:        params.UpdatedAt,
		deletedAt:        params.DeletedAt,
	}
}

// <-- Getters -->

func (e *Employee) ID() value_object.EmployeeID                     { return e.id }
func (e *Employee) UserID() value_object.UserID                     { return e.userID }
func (e *Employee) OrganizationID() value_object.OrganizationID     { return e.organizationID }
func (e *Employee) JobTitle() string                                { return e.jobTitle }
func (e *Employee) MembershipStatus() value_object.MembershipStatus { return e.membershipStatus }
func (e *Employee) EmploymentType() value_object.EmploymentType     { return e.employmentType }
func (e *Employee) CreatedAt() time.Time                            { return e.createdAt }
func (e *Employee) UpdatedAt() time.Time                            { return e.updatedAt }
func (e *Employee) DeletedAt() *time.Time                           { return e.deletedAt }

// <-- Helpers -->

func validateJobTitle(jobTitle string) (string, error) {
	jobTitle = strings.TrimSpace(jobTitle)

	length := len([]rune(jobTitle))

	if length < MinJobTitleLength {
		return "", ErrJobTitleTooShort
	}

	if length > MaxJobTitleLength {
		return "", ErrJobTitleTooLong
	}

	return jobTitle, nil
}
