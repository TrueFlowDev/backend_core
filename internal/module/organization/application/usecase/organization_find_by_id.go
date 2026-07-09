package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/value_object"
)

type FindOrganizationByIDInput struct {
	ID string
}

type FindOrganizationByIDUsecase struct {
	organizationRepository port.OrganizationRepository
}

func NewFindOrganizationByIDUsecase(
	organizationRepository port.OrganizationRepository,
) *FindOrganizationByIDUsecase {
	return &FindOrganizationByIDUsecase{
		organizationRepository: organizationRepository,
	}
}

func (u *FindOrganizationByIDUsecase) Execute(
	ctx context.Context,
	input FindOrganizationByIDInput,
) (*entity.Organization, error) {
	organizationID := value_object.NewOrganizationID(input.ID)

	organization, err := u.organizationRepository.FindByID(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
