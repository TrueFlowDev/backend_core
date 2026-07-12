package adapter

import (
	"context"

	employeeUsecase "github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
)

var _ port.EmployeeCreator = (*EmployeeCreator)(nil)

type EmployeeCreator struct {
	createEmployeeUsecase *employeeUsecase.CreateEmployeeUsecase
}

func NewEmployeeCreator(
	createEmployeeUsecase *employeeUsecase.CreateEmployeeUsecase,
) *EmployeeCreator {
	return &EmployeeCreator{createEmployeeUsecase: createEmployeeUsecase}
}

func (a *EmployeeCreator) Create(ctx context.Context, input port.EmployeeCreatorInput) (port.EmployeeCreatorOutput, error) {
	output, err := a.createEmployeeUsecase.Execute(ctx, employeeUsecase.CreateEmployeeInput{
		UserID:           input.UserID,
		OrganizationID:   input.OrganizationID,
		RoleID:           input.RoleID,
		JobTitle:         input.JobTitle,
		MembershipStatus: input.MembershipStatus,
		EmploymentType:   input.EmploymentType,
	})
	if err != nil {
		return port.EmployeeCreatorOutput{}, err
	}
	return port.EmployeeCreatorOutput{ID: output.ID}, nil
}
