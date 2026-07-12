package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/employee/infrastructure/model"
	"gorm.io/gorm"
)

func EmployeeModelToEntity(m *model.Employee) (*entity.Employee, error) {
	employeeID := valueobject.NewEmployeeID(m.ID)
	userID := valueobject.NewUserID(m.UserID)
	organizationID := valueobject.NewOrganizationID(m.OrganizationID)
	roleID := valueobject.NewRoleID(m.RoleID)

	membershipStatus, err := valueobject.ParseMembershipStatus(m.MembershipStatus)
	if err != nil {
		return nil, err
	}

	employmentType, err := valueobject.ParseEmploymentType(m.EmploymentType)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return entity.RestoreEmployee(entity.RestoreEmployeeParams{
		ID:               employeeID,
		UserID:           userID,
		OrganizationID:   organizationID,
		RoleID:           roleID,
		JobTitle:         m.JobTitle,
		MembershipStatus: membershipStatus,
		EmploymentType:   employmentType,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        deletedAt,
	}), nil
}

func EmployeeEntityToModel(e *entity.Employee) *model.Employee {
	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *e.DeletedAt(),
			Valid: true,
		}
	}

	return &model.Employee{
		ID:               e.ID().Value(),
		UserID:           e.UserID().Value(),
		OrganizationID:   e.OrganizationID().Value(),
		RoleID:           e.RoleID().Value(),
		JobTitle:         e.JobTitle(),
		MembershipStatus: e.MembershipStatus().Value(),
		EmploymentType:   e.EmploymentType().Value(),
		CreatedAt:        e.CreatedAt(),
		UpdatedAt:        e.UpdatedAt(),
		DeletedAt:        deletedAt,
	}
}
