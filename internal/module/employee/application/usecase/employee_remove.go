package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type RemoveEmployeeInput struct {
	ID               string
	OrganizationID   string
	RequestingUserID string
}

type RemoveEmployeeUsecase struct {
	employeeRepository port.EmployeeRepository
	roleFinder         port.RoleFinder
}

func NewRemoveEmployeeUsecase(
	employeeRepository port.EmployeeRepository,
	roleFinder port.RoleFinder,
) *RemoveEmployeeUsecase {
	return &RemoveEmployeeUsecase{
		employeeRepository: employeeRepository,
		roleFinder:         roleFinder,
	}
}

func (u *RemoveEmployeeUsecase) Execute(ctx context.Context, input RemoveEmployeeInput) error {
	employeeID := valueobject.NewEmployeeID(input.ID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	employee, err := u.employeeRepository.FindByID(ctx, employeeID, organizationID)
	if err != nil {
		return err
	}

	if employee.UserID().Value() == input.RequestingUserID {
		return entity.ErrCannotRemoveSelf
	}

	role, err := u.roleFinder.FindByID(ctx, employee.RoleID().Value(), input.OrganizationID)
	if err != nil {
		return err
	}
	if role.IsOwner {
		return entity.ErrCannotRemoveOwner
	}

	return u.employeeRepository.Delete(ctx, employeeID, organizationID)
}
