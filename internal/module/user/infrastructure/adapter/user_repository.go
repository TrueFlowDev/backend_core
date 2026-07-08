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
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"gorm.io/gorm"
)

var _ port.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	*database.BaseRepo
}

func NewUserRepository(base *database.BaseRepo) *UserRepository {
	return &UserRepository{BaseRepo: base}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	q := dao.Use(r.Executor(ctx))

	mappedUser := mapper.UserEntityToModel(user)
	if err := q.WithContext(ctx).User.Create(mappedUser); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return port.ErrUserAlreadyExists
		}
		return xerr.Wrap(err, port.ErrUserRepository.Code(), xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_create"))
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id value_object.UserID) (*entity.User, error) {
	q := dao.Use(r.Executor(ctx))

	model, err := q.WithContext(ctx).User.
		Where(q.User.ID.Eq(id.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrUserNotFound
		}
		return nil, xerr.Wrap(err, port.ErrUserRepository.Code(), xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_find_by_id"))
	}
	mappedUser, err := mapper.UserModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone value_object.Phone) (*entity.User, error) {
	q := dao.Use(r.Executor(ctx))

	model, err := q.WithContext(ctx).User.
		Where(q.User.Phone.Eq(phone.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrUserNotFound
		}
		return nil, xerr.Wrap(err, port.ErrUserRepository.Code(), xerr.WithDiagnostics(xerr.DiagnosticOperation, "user_find_by_phone"))
	}
	mappedUser, err := mapper.UserModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}
