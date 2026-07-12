package adapter

import (
	"context"

	authorizationUsecase "github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
)

var _ port.RoleCreator = (*RoleCreator)(nil)

type RoleCreator struct {
	createOwnerRoleUsecase *authorizationUsecase.CreateOwnerRoleUsecase
}

func NewRoleCreator(
	createOwnerRoleUsecase *authorizationUsecase.CreateOwnerRoleUsecase,
) *RoleCreator {
	return &RoleCreator{createOwnerRoleUsecase: createOwnerRoleUsecase}
}

func (a *RoleCreator) CreateOwnerRole(ctx context.Context, input port.RoleCreatorInput) (port.RoleCreatorOutput, error) {
	output, err := a.createOwnerRoleUsecase.Execute(ctx, authorizationUsecase.CreateOwnerRoleInput{
		OrganizationID: input.OrganizationID,
	})
	if err != nil {
		return port.RoleCreatorOutput{}, err
	}
	return port.RoleCreatorOutput{ID: output.ID}, nil
}
