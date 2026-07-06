package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
)

var (
	ErrUserNotFound      = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("user", xerr.ErrorReasonNotFound))
	ErrUserAlreadyExists = xerr.New(xerr.CodeAlreadyExists, xerr.WithMeta("user", xerr.ErrorReasonAlreadyExists))
)

type UserFinderByPhoneInput struct {
	Phone string
}

type UserFinderByPhoneOutput struct {
	ID             string
	HashedPassword string
}

type UserFinderByPhone interface {
	FindByPhone(ctx context.Context, input UserFinderByPhoneInput) (UserFinderByPhoneOutput, error)
}
