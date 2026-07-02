package port

import "github.com/Ali127Dev/xerr"

var (
	ErrFailedToGenerateOtp = xerr.New(xerr.CodeInternalError)
)

type OtpCodeGenerator interface {
	Generate() (string, error)
}
