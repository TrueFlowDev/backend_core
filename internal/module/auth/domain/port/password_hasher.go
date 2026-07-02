package port

import "github.com/Ali127Dev/xerr"

var (
	ErrInvalidPassword = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("password", "invalid_format"),
	)

	ErrPasswordMismatch = xerr.New(
		xerr.CodeUnauthorized,
		xerr.WithMeta("reason", "password_mismatch"),
	)

	ErrInvalidHash = xerr.New(
		xerr.CodeInternalError,
		xerr.WithMeta("reason", "corrupted_hash"),
	)
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Validate(password string, hashedPassword string) (bool, error)
}
