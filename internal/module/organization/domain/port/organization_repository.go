package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
)

var (
	ErrOrganizationNotFound   = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("organization", xerr.ErrorReasonNotFound))
	ErrOrganizationRepository = xerr.New(xerr.CodeDatabaseError)
)

type OrganizationRepository interface {
	Create(ctx context.Context, organization *entity.Organization) error

	FindByID(ctx context.Context, id valueobject.OrganizationID) (*entity.Organization, error)
	FindByIDs(ctx context.Context, ids []valueobject.OrganizationID) ([]*entity.Organization, error)
}
