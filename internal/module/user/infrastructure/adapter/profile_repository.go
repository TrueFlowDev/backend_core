package adapter

import (
	"context"
	"errors"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/dao"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/mapper"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ port.ProfileRepository = (*ProfileRepository)(nil)

type ProfileRepository struct {
	*database.BaseRepo
}

func NewProfileRepository(base *database.BaseRepo) *ProfileRepository {
	return &ProfileRepository{BaseRepo: base}
}

func (r *ProfileRepository) Save(ctx context.Context, profile *entity.Profile) error {
	q := dao.Use(r.Executor(ctx))

	mappedProfile := mapper.ProfileEntityToModel(profile)
	if err := q.WithContext(ctx).
		UsersProfile.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
			},
			DoUpdates: clause.AssignmentColumns(
				[]string{
					"email",
					"first_name",
					"last_name",
					"headline",
					"bio",
					"updated_at",
				},
			),
		}).
		Create(mappedProfile); err != nil {
		return xerr.Wrap(
			err, port.ErrProfileRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_profile_save"),
		)
	}

	return nil
}

func (r *ProfileRepository) FindByUserID(ctx context.Context, id valueobject.UserID) (*entity.Profile, error) {
	q := dao.Use(r.Executor(ctx))

	model, err := q.WithContext(ctx).UsersProfile.
		Where(q.UsersProfile.UserID.Eq(id.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrProfileNotFound
		}
		return nil, xerr.Wrap(
			err, port.ErrProfileRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_profile_find_by_user_id"),
		)
	}
	mappedProfile, err := mapper.ProfileModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedProfile, nil
}

func (r *ProfileRepository) FindByUserIDs(ctx context.Context, ids []valueobject.UserID) ([]*entity.Profile, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	q := dao.Use(r.Executor(ctx))

	rawIDs := make([]string, len(ids))
	for i, id := range ids {
		rawIDs[i] = id.Value()
	}

	models, err := q.WithContext(ctx).UsersProfile.
		Where(q.UsersProfile.UserID.In(rawIDs...)).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrProfileRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_profile_find_by_user_ids"))
	}

	profiles := make([]*entity.Profile, 0, len(models))
	for _, m := range models {
		p, err := mapper.ProfileModelToEntity(m)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}

	return profiles, nil
}
