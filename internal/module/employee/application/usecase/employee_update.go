package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type UpdateEmployeeInput struct {
	ID               string
	OrganizationID   string
	RequestingUserID string
	RoleID           string
	JobTitle         string
	MembershipStatus string
	EmploymentType   string
}

type UpdateEmployeeUsecase struct {
	employeeRepository port.EmployeeRepository
	roleFinder         port.RoleFinder
}

func NewUpdateEmployeeUsecase(
	employeeRepository port.EmployeeRepository,
	roleFinder port.RoleFinder,
) *UpdateEmployeeUsecase {
	return &UpdateEmployeeUsecase{
		employeeRepository: employeeRepository,
		roleFinder:         roleFinder,
	}
}

func (u *UpdateEmployeeUsecase) Execute(ctx context.Context, input UpdateEmployeeInput) error {
	employeeID := valueobject.NewEmployeeID(input.ID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	employee, err := u.employeeRepository.FindByID(ctx, employeeID, organizationID)
	if err != nil {
		return err
	}

	roleChanged := input.RoleID != employee.RoleID().Value()

	if roleChanged {
		if employee.UserID().Value() == input.RequestingUserID {
			return entity.ErrCannotChangeSelfRole
		}

		currentRole, err := u.roleFinder.FindByID(ctx, employee.RoleID().Value(), input.OrganizationID)
		if err != nil {
			return err
		}
		if currentRole.IsOwner {
			return entity.ErrOwnerRoleCannotBeReassigned
		}

		newRole, err := u.roleFinder.FindByID(ctx, input.RoleID, input.OrganizationID)
		if err != nil {
			return err
		}
		if newRole.IsOwner {
			return entity.ErrOwnerRoleCannotBeReassigned
		}
	}

	membershipStatus, err := valueobject.ParseMembershipStatus(input.MembershipStatus)
	if err != nil {
		return err
	}

	employmentType, err := valueobject.ParseEmploymentType(input.EmploymentType)
	if err != nil {
		return err
	}

	if err := employee.Update(input.JobTitle, valueobject.NewRoleID(input.RoleID), membershipStatus, employmentType); err != nil {
		return err
	}

	return u.employeeRepository.Update(ctx, employee)
}
