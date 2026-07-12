package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

type HasPermissionInput struct {
	UserID         string
	OrganizationID string
	Permission     string
}

type HasPermissionUsecase struct {
	employeeRoleFinder port.EmployeeRoleFinder
	roleRepository     port.RoleRepository
}

func NewHasPermissionUsecase(
	employeeRoleFinder port.EmployeeRoleFinder,
	roleRepository port.RoleRepository,
) *HasPermissionUsecase {
	return &HasPermissionUsecase{
		employeeRoleFinder: employeeRoleFinder,
		roleRepository:     roleRepository,
	}
}

func (u *HasPermissionUsecase) Execute(ctx context.Context, input HasPermissionInput) (bool, error) {
	roleFinderOutput, err := u.employeeRoleFinder.FindRoleID(ctx, input.UserID, input.OrganizationID)
	if err != nil {
		return false, err
	}

	roleID := valueobject.NewRoleID(roleFinderOutput.RoleID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	role, err := u.roleRepository.FindByID(ctx, roleID, organizationID)
	if err != nil {
		return false, err
	}

	if role.IsOwner() {
		return true, nil
	}

	requiredPermission, err := valueobject.ParsePermission(input.Permission)
	if err != nil {
		return false, err
	}

	for _, p := range role.Permissions() {
		if p.Value() == requiredPermission.Value() {
			return true, nil
		}
	}

	return false, nil
}
