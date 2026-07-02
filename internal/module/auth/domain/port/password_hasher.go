package port

import "github.com/Ali127Dev/xerr"

var (
	ErrInvalidPassword = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("password", xerr.ErrorReasonInvalidFormat),
	)

	ErrPasswordMismatch = xerr.New(
		xerr.CodeUnauthorized,
		xerr.WithMeta("password", xerr.ErrorReasonMismatch),
	)

	ErrInvalidHash = xerr.New(
		xerr.CodeInternalError,
		xerr.WithMeta("password", xerr.ErrorReasonCorrupted),
	)
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Validate(password string, hashedPassword string) (bool, error)
}
