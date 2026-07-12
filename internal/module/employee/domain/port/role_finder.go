package port

import "context"

type RoleFinderOutput struct {
	ID      string
	IsOwner bool
}

type RoleFinder interface {
	FindByID(ctx context.Context, roleID string, organizationID string) (RoleFinderOutput, error)
}
