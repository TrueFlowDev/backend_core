package adapter

import (
	"context"
	"errors"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/organization/infrastructure/dao"
	"github.com/TrueFlowDev/Backend/internal/module/organization/infrastructure/mapper"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"gorm.io/gorm"
)

var _ port.OrganizationRepository = (*OrganizationRepository)(nil)

type OrganizationRepository struct {
	*database.BaseRepo
}

func NewOrganizationRepository(base *database.BaseRepo) *OrganizationRepository {
	return &OrganizationRepository{BaseRepo: base}
}

func (o *OrganizationRepository) Create(ctx context.Context, organization *entity.Organization) error {
	q := dao.Use(o.Executor(ctx))

	mappedOrganization := mapper.OrganizationEntityToModel(organization)
	if err := q.WithContext(ctx).Organization.Create(mappedOrganization); err != nil {
		return xerr.Wrap(err, port.ErrOrganizationRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "organization_create"))
	}

	return nil
}

func (o *OrganizationRepository) Update(ctx context.Context, organization *entity.Organization) error {
	q := dao.Use(o.Executor(ctx))

	organizationModel := mapper.OrganizationEntityToModel(organization)

	result, err := q.WithContext(ctx).Organization.
		Where(q.Organization.ID.Eq(organizationModel.ID)).
		Updates(organizationModel)
	if err != nil {
		return xerr.Wrap(err, port.ErrOrganizationRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "organization_update"))
	}
	if result.RowsAffected == 0 {
		return port.ErrOrganizationNotFound
	}

	return nil
}

func (o *OrganizationRepository) FindByID(ctx context.Context, id valueobject.OrganizationID) (*entity.Organization, error) {
	q := dao.Use(o.Executor(ctx))

	model, err := q.WithContext(ctx).Organization.
		Where(q.Organization.ID.Eq(id.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrOrganizationNotFound
		}
		return nil, xerr.Wrap(err, port.ErrOrganizationRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "organization_find_by_id"))
	}
	mappedUser, err := mapper.OrganizationModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}

func (o *OrganizationRepository) FindByIDs(
	ctx context.Context, ids []valueobject.OrganizationID,
) ([]*entity.Organization, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	q := dao.Use(o.Executor(ctx))

	rawIDs := make([]string, len(ids))
	for i, id := range ids {
		rawIDs[i] = id.Value()
	}

	models, err := q.WithContext(ctx).Organization.
		Where(q.Organization.ID.In(rawIDs...)).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrOrganizationRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "organization_find_by_ids"))
	}

	organizations := make([]*entity.Organization, 0, len(models))
	for _, m := range models {
		org, err := mapper.OrganizationModelToEntity(m)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}

	return organizations, nil
}
