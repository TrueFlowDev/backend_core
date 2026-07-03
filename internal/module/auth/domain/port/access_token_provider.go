package port

import (
	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

var (
	ErrFailedToSignToken = xerr.New(xerr.CodeInternalError)
	InvalidToken         = xerr.New(xerr.CodeInvalidToken, xerr.WithMeta("token", xerr.ErrorReasonInvalidFormat))
)

type AccessTokenProvider interface {
	Generate(claims value_object.AccessTokenClaims) (string, error)
	Verify(token string) (value_object.AccessTokenClaims, error)
}
