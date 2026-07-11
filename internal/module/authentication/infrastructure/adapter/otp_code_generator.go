package adapter

import (
	"crypto/rand"
	"math/big"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
)

var _ port.OtpCodeGenerator = (*OtpCodeGenerator)(nil)

const digits = "0123456789"

type OtpCodeGenerator struct{ length int }

func NewOtpCodeGenerator(cfg *config.Config) *OtpCodeGenerator {
	return &OtpCodeGenerator{length: cfg.OTP.Length}
}

func (g *OtpCodeGenerator) Generate() (string, error) {
	code := make([]byte, g.length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", xerr.Wrap(
				err,
				port.ErrFailedToGenerateOtp.Code(),
				xerr.WithDiagnostics(xerr.DiagnosticOperation, "otp_generate"),
			)
		}
		code[i] = digits[n.Int64()]
	}
	return string(code), nil
}
