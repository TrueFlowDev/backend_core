package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
)

type FindOrganizationsByIDsInput struct {
	IDs []string
}

type FindOrganizationsByIDsUsecase struct {
	organizationRepository port.OrganizationRepository
}

func NewFindOrganizationsByIDsUsecase(
	organizationRepository port.OrganizationRepository,
) *FindOrganizationsByIDsUsecase {
	return &FindOrganizationsByIDsUsecase{
		organizationRepository: organizationRepository,
	}
}

func (u *FindOrganizationsByIDsUsecase) Execute(
	ctx context.Context, input FindOrganizationsByIDsInput,
) ([]*entity.Organization, error) {
	ids := make([]valueobject.OrganizationID, len(input.IDs))
	for i, raw := range input.IDs {
		ids[i] = valueobject.NewOrganizationID(raw)
	}

	return u.organizationRepository.FindByIDs(ctx, ids)
}
