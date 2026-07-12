package port

import (
	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/valueobject"
)

var (
	ErrFailedToSignToken = xerr.New(xerr.CodeInternalError)
	ErrTokenExpired      = xerr.New(xerr.CodeInvalidToken, xerr.WithMeta("token", xerr.ErrorReasonExpired))
	ErrInvalidToken      = xerr.New(xerr.CodeInvalidToken, xerr.WithMeta("token", xerr.ErrorReasonInvalidFormat))
)

type AccessTokenProvider interface {
	Generate(claims valueobject.AccessTokenClaims) (string, error)
	Verify(token string) (valueobject.AccessTokenClaims, error)
}
