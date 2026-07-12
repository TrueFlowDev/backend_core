package port

import "context"

type EmployeeRoleFinderOutput struct {
	RoleID string
}

type EmployeeRoleFinder interface {
	FindRoleID(ctx context.Context, userID string, organizationID string) (EmployeeRoleFinderOutput, error)
}
