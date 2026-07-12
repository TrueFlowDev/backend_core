package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/organization/infrastructure/model"
	"gorm.io/gorm"
)

func OrganizationModelToEntity(m *model.Organization) (*entity.Organization, error) {
	category, err := valueobject.ParseOrganizationCategory(m.Category)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}
	organization := entity.RestoreOrganization(entity.RestoreOrganizationParams{
		ID:        valueobject.NewOrganizationID(m.ID),
		Category:  category,
		Name:      m.Name,
		Active:    m.Active,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAt,
	})
	return organization, nil
}

func OrganizationEntityToModel(e *entity.Organization) *model.Organization {
	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{Time: *e.DeletedAt(), Valid: true}
	}

	return &model.Organization{
		ID:        e.ID().Value(),
		Category:  e.Category().Value(),
		Name:      e.Name(),
		Active:    e.Active(),
		CreatedAt: e.CreatedAt(),
		UpdatedAt: e.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
