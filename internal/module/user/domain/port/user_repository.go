package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

var (
	ErrUserNotFound      = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("user", xerr.ErrorReasonNotFound))
	ErrUserAlreadyExists = xerr.New(xerr.CodeAlreadyExists, xerr.WithMeta("user", xerr.ErrorReasonAlreadyExists))
	ErrUserRepository    = xerr.New(xerr.CodeDatabaseError)
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error

	FindByID(ctx context.Context, id value_object.UserID) (*entity.User, error)
	FindByPhone(ctx context.Context, phone value_object.Phone) (*entity.User, error)
}
