package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	sharedPort "github.com/TrueFlowDev/Backend/internal/shared/domain/port"
)

type DeleteRoleInput struct {
	ID             string
	OrganizationID string
}

type DeleteRoleUsecase struct {
	roleRepository           port.RoleRepository
	employeeRoleUsageChecker port.EmployeeRoleUsageChecker
	txManager                sharedPort.TxManager
}

func NewDeleteRoleUsecase(
	roleRepository port.RoleRepository,
	employeeRoleUsageChecker port.EmployeeRoleUsageChecker,
	txManager sharedPort.TxManager,
) *DeleteRoleUsecase {
	return &DeleteRoleUsecase{
		roleRepository:           roleRepository,
		employeeRoleUsageChecker: employeeRoleUsageChecker,
		txManager:                txManager,
	}
}

func (u *DeleteRoleUsecase) Execute(ctx context.Context, input DeleteRoleInput) error {
	roleID := valueobject.NewRoleID(input.ID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	role, err := u.roleRepository.FindByID(ctx, roleID, organizationID)
	if err != nil {
		return err
	}

	if role.IsOwner() {
		return entity.ErrCannotModifyOwnerRole
	}

	employeeCount, err := u.employeeRoleUsageChecker.CountActiveEmployeesByRole(ctx, input.ID)
	if err != nil {
		return err
	}
	if employeeCount > 0 {
		return port.ErrRoleInUse
	}

	return u.txManager.WithinTx(ctx, func(ctx context.Context) error {
		return u.roleRepository.Delete(ctx, roleID, organizationID)
	})
}
