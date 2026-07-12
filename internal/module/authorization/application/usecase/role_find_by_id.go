package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

type FindRoleByIDInput struct {
	ID string
}

type FindRoleByIDUsecase struct {
	roleRepository port.RoleRepository
}

func NewFindRoleByIDUsecase(
	roleRepository port.RoleRepository,
) *FindRoleByIDUsecase {
	return &FindRoleByIDUsecase{
		roleRepository: roleRepository,
	}
}

func (u *FindRoleByIDUsecase) Execute(ctx context.Context, input FindRoleByIDInput) (*entity.Role, error) {
	roleID := valueobject.NewRoleID(input.ID)

	role, err := u.roleRepository.FindByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return role, nil
}
