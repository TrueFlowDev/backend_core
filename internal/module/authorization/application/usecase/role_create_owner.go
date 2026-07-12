package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

type CreateOwnerRoleInput struct {
	OrganizationID string
}

type CreateOwnerRoleOutput struct {
	ID string
}

type CreateOwnerRoleUsecase struct {
	roleIDGenerator port.RoleIDGenerator
	roleRepository  port.RoleRepository
}

func NewCreateOwnerRoleUsecase(
	roleIDGenerator port.RoleIDGenerator,
	roleRepository port.RoleRepository,
) *CreateOwnerRoleUsecase {
	return &CreateOwnerRoleUsecase{
		roleIDGenerator: roleIDGenerator,
		roleRepository:  roleRepository,
	}
}

func (u *CreateOwnerRoleUsecase) Execute(ctx context.Context, input CreateOwnerRoleInput) (CreateOwnerRoleOutput, error) {
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	newRoleID := u.roleIDGenerator.Generate()

	ownerRole := entity.NewOwnerRole(newRoleID, organizationID)

	if err := u.roleRepository.Create(ctx, ownerRole); err != nil {
		return CreateOwnerRoleOutput{}, err
	}

	return CreateOwnerRoleOutput{ID: newRoleID.Value()}, nil
}
