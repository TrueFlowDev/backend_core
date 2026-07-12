package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

type ListRolesInput struct {
	OrganizationID string
}

type ListRolesUsecase struct {
	roleRepository port.RoleRepository
}

func NewListRolesUsecase(
	roleRepository port.RoleRepository,
) *ListRolesUsecase {
	return &ListRolesUsecase{roleRepository: roleRepository}
}

func (u *ListRolesUsecase) Execute(ctx context.Context, input ListRolesInput) ([]*entity.Role, error) {
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	return u.roleRepository.ListByOrganizationID(ctx, organizationID)
}
