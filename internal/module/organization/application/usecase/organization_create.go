package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
)

type CreateOrganizationInput struct {
	Category string
	Name     string
}

type CreateOrganizationOutput struct {
	ID string
}

type CreateOrganizationUsecase struct {
	organizationIDGenerator port.OrganizationIDGenerator
	organizationRepository  port.OrganizationRepository
}

func NewCreateOrganizationUsecase(
	organizationIDGenerator port.OrganizationIDGenerator,
	organizationRepository port.OrganizationRepository,
) *CreateOrganizationUsecase {
	return &CreateOrganizationUsecase{
		organizationIDGenerator: organizationIDGenerator,
		organizationRepository:  organizationRepository,
	}
}

func (u *CreateOrganizationUsecase) Execute(ctx context.Context, input CreateOrganizationInput) (CreateOrganizationOutput, error) {
	newOrganizationID := u.organizationIDGenerator.Generate()

	newOrganizationCategory, err := valueobject.NewOrganizationCategory(input.Category)
	if err != nil {
		return CreateOrganizationOutput{}, err
	}

	newOrganization, err := entity.NewOrganization(newOrganizationID, input.Name, newOrganizationCategory)
	if err != nil {
		return CreateOrganizationOutput{}, err
	}

	if err := u.organizationRepository.Create(ctx, newOrganization); err != nil {
		return CreateOrganizationOutput{}, err
	}

	return CreateOrganizationOutput{ID: newOrganizationID.Value()}, nil
}
