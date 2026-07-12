package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	employeeUsecase "github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
)

var _ port.EmployeeRoleUsageChecker = (*EmployeeRoleUsageChecker)(nil)

type EmployeeRoleUsageChecker struct {
	countActiveEmployeesByRoleUsecase *employeeUsecase.CountActiveEmployeesByRoleUsecase
}

func NewEmployeeRoleUsageChecker(
	countActiveEmployeesByRoleUsecase *employeeUsecase.CountActiveEmployeesByRoleUsecase,
) *EmployeeRoleUsageChecker {
	return &EmployeeRoleUsageChecker{countActiveEmployeesByRoleUsecase: countActiveEmployeesByRoleUsecase}
}

func (a *EmployeeRoleUsageChecker) CountActiveEmployeesByRole(ctx context.Context, roleID string) (int64, error) {
	return a.countActiveEmployeesByRoleUsecase.Execute(ctx, employeeUsecase.CountActiveEmployeesByRoleInput{
		RoleID: roleID,
	})
}
