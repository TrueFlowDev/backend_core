package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	employeeUsecase "github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
)

var _ port.EmployeeRoleFinder = (*EmployeeRoleFinder)(nil)

type EmployeeRoleFinder struct {
	findEmployeeRoleUsecase *employeeUsecase.FindEmployeeRoleUsecase
}

func NewEmployeeRoleFinder(
	findEmployeeRoleUsecase *employeeUsecase.FindEmployeeRoleUsecase,
) *EmployeeRoleFinder {
	return &EmployeeRoleFinder{findEmployeeRoleUsecase: findEmployeeRoleUsecase}
}

func (a *EmployeeRoleFinder) FindRoleID(ctx context.Context, userID string, organizationID string) (port.EmployeeRoleFinderOutput, error) {
	output, err := a.findEmployeeRoleUsecase.Execute(ctx, employeeUsecase.FindEmployeeRoleInput{
		UserID:         userID,
		OrganizationID: organizationID,
	})
	if err != nil {
		return port.EmployeeRoleFinderOutput{}, err
	}

	return port.EmployeeRoleFinderOutput{RoleID: output.RoleID}, nil
}
