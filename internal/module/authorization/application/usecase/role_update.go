package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	sharedPort "github.com/TrueFlowDev/Backend/internal/shared/domain/port"
)

type UpdateRoleInput struct {
	ID             string
	OrganizationID string
	Title          string
	Permissions    []string
}

type UpdateRoleUsecase struct {
	roleRepository port.RoleRepository
	txManager      sharedPort.TxManager
}

func NewUpdateRoleUsecase(
	roleRepository port.RoleRepository,
	txManager sharedPort.TxManager,
) *UpdateRoleUsecase {
	return &UpdateRoleUsecase{
		roleRepository: roleRepository,
		txManager:      txManager,
	}
}

func (u *UpdateRoleUsecase) Execute(ctx context.Context, input UpdateRoleInput) error {
	roleID := valueobject.NewRoleID(input.ID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	permissions := make([]valueobject.Permission, 0, len(input.Permissions))
	for _, raw := range input.Permissions {
		p, err := valueobject.ParsePermission(raw)
		if err != nil {
			return err
		}
		permissions = append(permissions, p)
	}

	return u.txManager.WithinTx(ctx, func(ctx context.Context) error {
		role, err := u.roleRepository.FindByID(ctx, roleID, organizationID)
		if err != nil {
			return err
		}

		if err := role.Update(input.Title, permissions); err != nil {
			return err
		}

		return u.roleRepository.Update(ctx, role)
	})
}
