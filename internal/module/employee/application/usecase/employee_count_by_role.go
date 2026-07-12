package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type CountActiveEmployeesByRoleInput struct {
	RoleID string
}

type CountActiveEmployeesByRoleUsecase struct {
	employeeRepository port.EmployeeRepository
}

func NewCountActiveEmployeesByRoleUsecase(
	employeeRepository port.EmployeeRepository,
) *CountActiveEmployeesByRoleUsecase {
	return &CountActiveEmployeesByRoleUsecase{employeeRepository: employeeRepository}
}

func (u *CountActiveEmployeesByRoleUsecase) Execute(ctx context.Context, input CountActiveEmployeesByRoleInput) (int64, error) {
	roleID := valueobject.NewRoleID(input.RoleID)
	return u.employeeRepository.CountActiveByRoleID(ctx, roleID)
}
