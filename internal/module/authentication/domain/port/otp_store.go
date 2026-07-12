package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/valueobject"
)

var (
	ErrOTPNotFound = xerr.New(xerr.CodeUnauthorized, xerr.WithMeta("otp", xerr.ErrorReasonNotFound))
	ErrOTPStore    = xerr.New(xerr.CodeDatabaseError)
)

type OTPStore interface {
	Set(ctx context.Context, key valueobject.Phone, value entity.OTP) error
	Get(ctx context.Context, key valueobject.Phone) (entity.OTP, error)
	Delete(ctx context.Context, key valueobject.Phone) error
}
