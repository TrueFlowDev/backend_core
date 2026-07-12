package adapter

import (
	"context"
	"errors"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/dao"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/mapper"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/model"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"gorm.io/gorm"
)

var _ port.RoleRepository = (*RoleRepository)(nil)

type RoleRepository struct {
	*database.BaseRepo
}

func NewRoleRepository(base *database.BaseRepo) *RoleRepository {
	return &RoleRepository{BaseRepo: base}
}

func (r *RoleRepository) Create(ctx context.Context, role *entity.Role) error {
	q := dao.Use(r.Executor(ctx))

	mappedRole, mappedPermissions := mapper.RoleEntityToModel(role)
	if err := q.WithContext(ctx).Role.Create(mappedRole); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return port.ErrRoleAlreadyExists
		}
		return xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_create"))
	}
	if err := q.WithContext(ctx).RolePermission.Create(mappedPermissions...); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return port.ErrRoleAlreadyExists
		}
		return xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_permissions_create"))
	}

	return nil
}

func (r *RoleRepository) FindByID(
	ctx context.Context, id valueobject.RoleID,
	organizationID valueobject.OrganizationID,
) (*entity.Role, error) {
	q := dao.Use(r.Executor(ctx))

	roleModel, err := q.WithContext(ctx).Role.
		Where(
			q.Role.ID.Eq(id.Value()),
			q.Role.OrganizationID.Eq(organizationID.Value()),
		).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrRoleNotFound
		}
		return nil, xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_find_by_id"))
	}

	permissionModels, err := q.WithContext(ctx).RolePermission.
		Where(q.RolePermission.RoleID.Eq(id.Value())).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_permissions_find_by_role_id"))
	}

	return mapper.RoleModelToEntity(roleModel, permissionModels)
}

func (r *RoleRepository) ListByOrganizationID(
	ctx context.Context, organizationID valueobject.OrganizationID,
) ([]*entity.Role, error) {
	q := dao.Use(r.Executor(ctx))

	roleModels, err := q.WithContext(ctx).Role.
		Where(q.Role.OrganizationID.Eq(organizationID.Value())).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_list_by_organization_id"))
	}

	if len(roleModels) == 0 {
		return []*entity.Role{}, nil
	}

	roleIDs := make([]string, len(roleModels))
	for i, m := range roleModels {
		roleIDs[i] = m.ID
	}

	permissionModels, err := q.WithContext(ctx).RolePermission.
		Where(q.RolePermission.RoleID.In(roleIDs...)).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrRoleRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "role_permissions_list_by_role_ids"))
	}

	permissionsByRoleID := make(map[string][]*model.RolePermission, len(roleModels))
	for _, pm := range permissionModels {
		permissionsByRoleID[pm.RoleID] = append(permissionsByRoleID[pm.RoleID], pm)
	}

	roles := make([]*entity.Role, 0, len(roleModels))
	for _, m := range roleModels {
		role, err := mapper.RoleModelToEntity(m, permissionsByRoleID[m.ID])
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
