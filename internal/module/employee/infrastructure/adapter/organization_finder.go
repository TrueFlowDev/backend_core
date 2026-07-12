package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	organizationUsecase "github.com/TrueFlowDev/Backend/internal/module/organization/application/usecase"
)

var _ port.OrganizationFinder = (*OrganizationFinder)(nil)

type OrganizationFinder struct {
	findOrganizationsByIDsUsecase *organizationUsecase.FindOrganizationsByIDsUsecase
}

func NewOrganizationFinder(
	findOrganizationsByIDsUsecase *organizationUsecase.FindOrganizationsByIDsUsecase,
) *OrganizationFinder {
	return &OrganizationFinder{findOrganizationsByIDsUsecase: findOrganizationsByIDsUsecase}
}

func (a *OrganizationFinder) FindByIDs(ctx context.Context, organizationIDs []string) ([]port.OrganizationFinderOutput, error) {
	organizations, err := a.findOrganizationsByIDsUsecase.Execute(ctx, organizationUsecase.FindOrganizationsByIDsInput{
		IDs: organizationIDs,
	})
	if err != nil {
		return nil, err
	}

	result := make([]port.OrganizationFinderOutput, len(organizations))
	for i, org := range organizations {
		result[i] = port.OrganizationFinderOutput{
			ID:       org.ID().Value(),
			Name:     org.Name(),
			Category: org.Category().Value(),
		}
	}

	return result, nil
}
