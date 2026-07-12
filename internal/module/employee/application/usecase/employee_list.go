package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type ListEmployeesInput struct {
	OrganizationID string
}

type EmployeeWithUserOutput struct {
	ID               string
	UserID           string
	RoleID           string
	JobTitle         string
	MembershipStatus string
	EmploymentType   string
	FirstName        string
	LastName         string
	Phone            string
}

type ListEmployeesUsecase struct {
	employeeRepository port.EmployeeRepository
	userFinder         port.UserFinder
}

func NewListEmployeesUsecase(
	employeeRepository port.EmployeeRepository,
	userFinder port.UserFinder,
) *ListEmployeesUsecase {
	return &ListEmployeesUsecase{
		employeeRepository: employeeRepository,
		userFinder:         userFinder,
	}
}

func (u *ListEmployeesUsecase) Execute(ctx context.Context, input ListEmployeesInput) ([]EmployeeWithUserOutput, error) {
	organizationID := valueobject.NewOrganizationID(input.OrganizationID)

	employees, err := u.employeeRepository.ListByOrganizationID(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	if len(employees) == 0 {
		return []EmployeeWithUserOutput{}, nil
	}

	userIDs := make([]string, len(employees))
	for i, e := range employees {
		userIDs[i] = e.UserID().Value()
	}

	users, err := u.userFinder.FindByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	userByID := make(map[string]port.UserFinderOutput, len(users))
	for _, user := range users {
		userByID[user.ID] = user
	}

	result := make([]EmployeeWithUserOutput, 0, len(employees))
	for _, e := range employees {
		user, ok := userByID[e.UserID().Value()]
		if !ok {
			continue
		}

		result = append(result, EmployeeWithUserOutput{
			ID:               e.ID().Value(),
			UserID:           e.UserID().Value(),
			RoleID:           e.RoleID().Value(),
			JobTitle:         e.JobTitle(),
			MembershipStatus: e.MembershipStatus().Value(),
			EmploymentType:   e.EmploymentType().Value(),
			FirstName:        user.FirstName,
			LastName:         user.LastName,
			Phone:            user.Phone,
		})
	}

	return result, nil
}
