package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
)

type UpdateOrganizationInput struct {
	ID       string
	Name     string
	Category string
}

type UpdateOrganizationUsecase struct {
	organizationRepository port.OrganizationRepository
}

func NewUpdateOrganizationUsecase(
	organizationRepository port.OrganizationRepository,
) *UpdateOrganizationUsecase {
	return &UpdateOrganizationUsecase{organizationRepository: organizationRepository}
}

func (u *UpdateOrganizationUsecase) Execute(ctx context.Context, input UpdateOrganizationInput) error {
	organizationID := valueobject.NewOrganizationID(input.ID)

	organization, err := u.organizationRepository.FindByID(ctx, organizationID)
	if err != nil {
		return err
	}

	category, err := valueobject.ParseOrganizationCategory(input.Category)
	if err != nil {
		return err
	}

	if err := organization.Update(input.Name, category); err != nil {
		return err
	}

	return u.organizationRepository.Update(ctx, organization)
}
