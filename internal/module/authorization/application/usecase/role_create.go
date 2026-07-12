package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	sharedPort "github.com/TrueFlowDev/Backend/internal/shared/domain/port"
)

type CreateRoleInput struct {
	OrganizationID string
	Title          string
	Permissions    []string
}

type CreateRoleOutput struct {
	ID string
}

type CreateRoleUsecase struct {
	roleIDGenerator port.RoleIDGenerator
	roleRepository  port.RoleRepository
	txManager       sharedPort.TxManager
}

func NewCreateRoleUsecase(
	roleIDGenerator port.RoleIDGenerator,
	roleRepository port.RoleRepository,
	txManager sharedPort.TxManager,
) *CreateRoleUsecase {
	return &CreateRoleUsecase{
		roleIDGenerator: roleIDGenerator,
		roleRepository:  roleRepository,
		txManager:       txManager,
	}
}

func (u *CreateRoleUsecase) Execute(ctx context.Context, input CreateRoleInput) (CreateRoleOutput, error) {
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	permissions := make([]valueobject.Permission, 0, len(input.Permissions))
	for _, raw := range input.Permissions {
		p, err := valueobject.ParsePermission(raw)
		if err != nil {
			return CreateRoleOutput{}, err
		}
		permissions = append(permissions, p)
	}

	newRoleID := u.roleIDGenerator.Generate()

	newRole, err := entity.NewRole(newRoleID, organizationID, input.Title, permissions)
	if err != nil {
		return CreateRoleOutput{}, err
	}

	if err := u.txManager.WithinTx(ctx, func(ctx context.Context) error {
		return u.roleRepository.Create(ctx, newRole)
	}); err != nil {
		return CreateRoleOutput{}, err
	}

	return CreateRoleOutput{ID: newRoleID.Value()}, nil
}
