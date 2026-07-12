package port

import "context"

type RoleCreatorInput struct {
	OrganizationID string
}

type RoleCreatorOutput struct {
	ID string
}

type RoleCreator interface {
	CreateOwnerRole(ctx context.Context, input RoleCreatorInput) (RoleCreatorOutput, error)
}
