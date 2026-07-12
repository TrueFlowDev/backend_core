package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
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
	id             valueobject.EmployeeID
	userID         valueobject.UserID
	organizationID valueobject.OrganizationID
	roleID         valueobject.RoleID

	jobTitle         string
	membershipStatus valueobject.MembershipStatus
	employmentType   valueobject.EmploymentType

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreEmployeeParams struct {
	ID             valueobject.EmployeeID
	UserID         valueobject.UserID
	OrganizationID valueobject.OrganizationID
	RoleID         valueobject.RoleID

	JobTitle         string
	MembershipStatus valueobject.MembershipStatus
	EmploymentType   valueobject.EmploymentType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewEmployee(
	id valueobject.EmployeeID,
	userID valueobject.UserID,
	organizationID valueobject.OrganizationID,
	roleID valueobject.RoleID,
	jobTitle string,
	membershipStatus valueobject.MembershipStatus,
	employmentType valueobject.EmploymentType,
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
		roleID:           roleID,
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
		roleID:           params.RoleID,
		jobTitle:         params.JobTitle,
		membershipStatus: params.MembershipStatus,
		employmentType:   params.EmploymentType,
		createdAt:        params.CreatedAt,
		updatedAt:        params.UpdatedAt,
		deletedAt:        params.DeletedAt,
	}
}

// <-- Getters -->

func (e *Employee) ID() valueobject.EmployeeID                     { return e.id }
func (e *Employee) UserID() valueobject.UserID                     { return e.userID }
func (e *Employee) OrganizationID() valueobject.OrganizationID     { return e.organizationID }
func (e *Employee) RoleID() valueobject.RoleID                     { return e.roleID }
func (e *Employee) JobTitle() string                               { return e.jobTitle }
func (e *Employee) MembershipStatus() valueobject.MembershipStatus { return e.membershipStatus }
func (e *Employee) EmploymentType() valueobject.EmploymentType     { return e.employmentType }
func (e *Employee) CreatedAt() time.Time                           { return e.createdAt }
func (e *Employee) UpdatedAt() time.Time                           { return e.updatedAt }
func (e *Employee) DeletedAt() *time.Time                          { return e.deletedAt }

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
