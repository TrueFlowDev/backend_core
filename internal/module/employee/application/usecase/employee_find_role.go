package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type FindEmployeeRoleInput struct {
	UserID         string
	OrganizationID string
}

type FindEmployeeRoleOutput struct {
	RoleID string
}

type FindEmployeeRoleUsecase struct {
	employeeRepository port.EmployeeRepository
}

func NewFindEmployeeRoleUsecase(
	employeeRepository port.EmployeeRepository,
) *FindEmployeeRoleUsecase {
	return &FindEmployeeRoleUsecase{employeeRepository: employeeRepository}
}

func (u *FindEmployeeRoleUsecase) Execute(ctx context.Context, input FindEmployeeRoleInput) (FindEmployeeRoleOutput, error) {
	userID := valueobject.NewUserID(input.UserID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	employee, err := u.employeeRepository.FindByUserIDAndOrganizationID(ctx, userID, organizationID)
	if err != nil {
		return FindEmployeeRoleOutput{}, err
	}

	return FindEmployeeRoleOutput{RoleID: employee.RoleID().Value()}, nil
}
