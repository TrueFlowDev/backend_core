package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/model"
	"gorm.io/gorm"
)

func RoleModelToEntity(m *model.Role, permissionModels []*model.RolePermission) (*entity.Role, error) {
	roleID := valueobject.NewRoleID(m.ID)
	organizationID := valueobject.NewOrganizationID(m.OrganizationID)

	permissions := make([]valueobject.Permission, 0, len(permissionModels))
	for _, pm := range permissionModels {
		p, err := valueobject.ParsePermission(pm.PermissionValue)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return entity.RestoreRole(entity.RestoreRoleParams{
		ID:             roleID,
		OrganizationID: organizationID,
		Title:          m.Title,
		Permissions:    permissions,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		DeletedAt:      deletedAt,
	}), nil
}

func RoleEntityToModel(e *entity.Role) (*model.Role, []*model.RolePermission) {
	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *e.DeletedAt(),
			Valid: true,
		}
	}

	roleModel := &model.Role{
		ID:             e.ID().Value(),
		OrganizationID: e.OrganizationID().Value(),
		Title:          e.Title(),
		CreatedAt:      e.CreatedAt(),
		UpdatedAt:      e.UpdatedAt(),
		DeletedAt:      deletedAt,
	}

	permissionModels := make([]*model.RolePermission, len(e.Permissions()))
	for i, p := range e.Permissions() {
		permissionModels[i] = &model.RolePermission{
			RoleID:          e.ID().Value(),
			PermissionValue: p.Value(),
		}
	}

	return roleModel, permissionModels
}
