package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

type ListMyDashboardsInput struct {
	UserID string
}

type DashboardOutput struct {
	EmployeeID       string
	OrganizationID   string
	OrganizationName string
	RoleID           string
	JobTitle         string
	EmploymentType   string
}

type ListMyDashboardsUsecase struct {
	employeeRepository port.EmployeeRepository
	organizationFinder port.OrganizationFinder
}

func NewListMyDashboardsUsecase(
	employeeRepository port.EmployeeRepository,
	organizationFinder port.OrganizationFinder,
) *ListMyDashboardsUsecase {
	return &ListMyDashboardsUsecase{
		employeeRepository: employeeRepository,
		organizationFinder: organizationFinder,
	}
}

func (u *ListMyDashboardsUsecase) Execute(ctx context.Context, input ListMyDashboardsInput) ([]DashboardOutput, error) {
	userID := valueobject.NewUserID(input.UserID)

	employees, err := u.employeeRepository.ListActiveByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(employees) == 0 {
		return []DashboardOutput{}, nil
	}

	organizationIDs := make([]string, len(employees))
	for i, e := range employees {
		organizationIDs[i] = e.OrganizationID().Value()
	}

	organizations, err := u.organizationFinder.FindByIDs(ctx, organizationIDs)
	if err != nil {
		return nil, err
	}

	organizationByID := make(map[string]port.OrganizationFinderOutput, len(organizations))
	for _, org := range organizations {
		organizationByID[org.ID] = org
	}

	result := make([]DashboardOutput, 0, len(employees))
	for _, e := range employees {
		organization, ok := organizationByID[e.OrganizationID().Value()]
		if !ok {
			continue
		}

		result = append(result, DashboardOutput{
			EmployeeID:       e.ID().Value(),
			OrganizationID:   organization.ID,
			OrganizationName: organization.Name,
			RoleID:           e.RoleID().Value(),
			JobTitle:         e.JobTitle(),
			EmploymentType:   e.EmploymentType().Value(),
		})
	}

	return result, nil
}
