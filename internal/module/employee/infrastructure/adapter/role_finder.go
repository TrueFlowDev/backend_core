package adapter

import (
	"context"

	authorizationUsecase "github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
)

var _ port.RoleFinder = (*RoleFinder)(nil)

type RoleFinder struct {
	findRoleByIDUsecase *authorizationUsecase.FindRoleByIDUsecase
}

func NewRoleFinder(
	findRoleByIDUsecase *authorizationUsecase.FindRoleByIDUsecase,
) *RoleFinder {
	return &RoleFinder{findRoleByIDUsecase: findRoleByIDUsecase}
}

func (a *RoleFinder) FindByID(ctx context.Context, roleID string, organizationID string) (port.RoleFinderOutput, error) {
	role, err := a.findRoleByIDUsecase.Execute(ctx, authorizationUsecase.FindRoleByIDInput{
		ID:             roleID,
		OrganizationID: organizationID,
	})
	if err != nil {
		return port.RoleFinderOutput{}, err
	}

	return port.RoleFinderOutput{ID: role.ID().Value(), IsOwner: role.IsOwner()}, nil
}
