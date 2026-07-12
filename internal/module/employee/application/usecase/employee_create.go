package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type CreateEmployeeInput struct {
	UserID           string
	OrganizationID   string
	RoleID           string
	JobTitle         string
	MembershipStatus string
	EmploymentType   string
}

type CreateEmployeeOutput struct {
	ID string
}

type CreateEmployeeUsecase struct {
	employeeIDGenerator port.EmployeeIDGenerator
	employeeRepository  port.EmployeeRepository
}

func NewCreateEmployeeUsecase(
	employeeIDGenerator port.EmployeeIDGenerator,
	employeeRepository port.EmployeeRepository,
) *CreateEmployeeUsecase {
	return &CreateEmployeeUsecase{
		employeeIDGenerator: employeeIDGenerator,
		employeeRepository:  employeeRepository,
	}
}

func (u *CreateEmployeeUsecase) Execute(ctx context.Context, input CreateEmployeeInput) (CreateEmployeeOutput, error) {
	userID := valueobject.NewUserID(input.UserID)
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)
	roleID := valueobject.NewRoleID(input.RoleID)

	membershipStatus, err := valueobject.ParseMembershipStatus(input.MembershipStatus)
	if err != nil {
		return CreateEmployeeOutput{}, err
	}

	employmentType, err := valueobject.ParseEmploymentType(input.EmploymentType)
	if err != nil {
		return CreateEmployeeOutput{}, err
	}

	newEmployeeID := u.employeeIDGenerator.Generate()

	newEmployee, err := entity.NewEmployee(
		newEmployeeID,
		userID,
		organizationID,
		roleID,
		input.JobTitle,
		membershipStatus,
		employmentType,
	)
	if err != nil {
		return CreateEmployeeOutput{}, err
	}

	if err := u.employeeRepository.Create(ctx, newEmployee); err != nil {
		return CreateEmployeeOutput{}, err
	}

	return CreateEmployeeOutput{ID: newEmployeeID.Value()}, nil
}
