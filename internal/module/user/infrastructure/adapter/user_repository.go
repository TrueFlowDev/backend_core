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
)

type UserRepository struct {
	db  *gorm.DB
	dao *dao.Query
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db, dao: dao.Use(db)}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	mappedUser := mapper.UserEntityToModel(user)
	if err := r.dao.WithContext(ctx).User.Create(mappedUser); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return port.ErrUserAlreadyExists
		}
		return xerr.Wrap(err, port.ErrUserRepository.Code())
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id value_object.UserID) (*entity.User, error) {
	model, err := r.dao.WithContext(ctx).User.
		Where(r.dao.User.ID.Eq(id.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrUserNotFound
		}
		return nil, xerr.Wrap(err, port.ErrUserRepository.Code())
	}
	mappedUser, err := mapper.UserModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone value_object.Phone) (*entity.User, error) {
	model, err := r.dao.WithContext(ctx).User.
		Where(r.dao.User.Phone.Eq(phone.Value())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrUserNotFound
		}
		return nil, xerr.Wrap(err, port.ErrUserRepository.Code())
	}
	mappedUser, err := mapper.UserModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}
