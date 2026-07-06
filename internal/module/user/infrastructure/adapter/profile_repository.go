package adapter

import (
	"context"
	"errors"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/dao"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/mapper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ port.ProfileRepository = (*ProfileRepository)(nil)

type ProfileRepository struct {
	db  *gorm.DB
	dao *dao.Query
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db, dao: dao.Use(db)}
}

func (r *ProfileRepository) Save(ctx context.Context, profile *entity.Profile) error {
	mappedProfile := mapper.ProfileEntityToModel(profile)
	if err := r.dao.WithContext(ctx).
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
			err, port.ErrUserRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_profile_create"),
		)
	}

	return nil
}

func (r *ProfileRepository) FindByUserID(ctx context.Context, id value_object.UserID) (*entity.Profile, error) {
	model, err := r.dao.WithContext(ctx).UsersProfile.
		Where(r.dao.UsersProfile.UserID.Eq(id.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrUserNotFound
		}
		return nil, xerr.Wrap(
			err, port.ErrProfileNotFound.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_profile_find_by_user_id"),
		)
	}
	mappedProfile, err := mapper.ProfileModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedProfile, nil
}
