package port

import (
	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/value_object"
)

var (
	ErrFailedToSignToken = xerr.New(xerr.CodeInternalError)
	ErrTokenExpired      = xerr.New(xerr.CodeInvalidToken, xerr.WithMeta("token", xerr.ErrorReasonExpired))
	ErrInvalidToken      = xerr.New(xerr.CodeInvalidToken, xerr.WithMeta("token", xerr.ErrorReasonInvalidFormat))
)

type AccessTokenProvider interface {
	Generate(claims value_object.AccessTokenClaims) (string, error)
	Verify(token string) (value_object.AccessTokenClaims, error)
}
