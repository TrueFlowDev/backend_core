package port

import "context"

type EmployeeCreatorInput struct {
	UserID           string
	OrganizationID   string
	RoleID           string
	JobTitle         string
	MembershipStatus string
	EmploymentType   string
}

type EmployeeCreatorOutput struct {
	ID string
}

type EmployeeCreator interface {
	Create(ctx context.Context, input EmployeeCreatorInput) (EmployeeCreatorOutput, error)
}
