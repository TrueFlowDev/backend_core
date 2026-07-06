package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

var (
	ErrProfileNotFound   = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("profile", xerr.ErrorReasonNotFound))
	ErrProfileRepository = xerr.New(xerr.CodeDatabaseError)
)

type ProfileRepository interface {
	Save(ctx context.Context, profile *entity.Profile) error

	FindByUserID(ctx context.Context, id value_object.UserID) (*entity.Profile, error)
}
