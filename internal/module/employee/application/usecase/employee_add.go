package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type AddEmployeeInput struct {
	OrganizationID   string
	Phone            string
	RoleID           string
	JobTitle         string
	MembershipStatus string
	EmploymentType   string
}

type AddEmployeeOutput struct {
	ID string
}

type AddEmployeeUsecase struct {
	employeeIDGenerator port.EmployeeIDGenerator
	employeeRepository  port.EmployeeRepository
	userFinder          port.UserFinder
	roleFinder          port.RoleFinder
}

func NewAddEmployeeUsecase(
	employeeIDGenerator port.EmployeeIDGenerator,
	employeeRepository port.EmployeeRepository,
	userFinder port.UserFinder,
	roleFinder port.RoleFinder,
) *AddEmployeeUsecase {
	return &AddEmployeeUsecase{
		employeeIDGenerator: employeeIDGenerator,
		employeeRepository:  employeeRepository,
		userFinder:          userFinder,
		roleFinder:          roleFinder,
	}
}

func (u *AddEmployeeUsecase) Execute(ctx context.Context, input AddEmployeeInput) (AddEmployeeOutput, error) {
	userOutput, err := u.userFinder.FindByPhone(ctx, input.Phone)
	if err != nil {
		return AddEmployeeOutput{}, err
	}

	if _, err := u.roleFinder.FindByID(ctx, input.RoleID, input.OrganizationID); err != nil {
		return AddEmployeeOutput{}, err
	}

	membershipStatus, err := valueobject.ParseMembershipStatus(input.MembershipStatus)
	if err != nil {
		return AddEmployeeOutput{}, err
	}

	employmentType, err := valueobject.ParseEmploymentType(input.EmploymentType)
	if err != nil {
		return AddEmployeeOutput{}, err
	}

	newEmployeeID := u.employeeIDGenerator.Generate()

	newEmployee, err := entity.NewEmployee(
		newEmployeeID,
		valueobject.NewUserID(userOutput.ID),
		valueobject.NewOrganizationID(input.OrganizationID),
		valueobject.NewRoleID(input.RoleID),
		input.JobTitle,
		membershipStatus,
		employmentType,
	)
	if err != nil {
		return AddEmployeeOutput{}, err
	}

	if err := u.employeeRepository.Create(ctx, newEmployee); err != nil {
		return AddEmployeeOutput{}, err
	}

	return AddEmployeeOutput{ID: newEmployeeID.Value()}, nil
}
