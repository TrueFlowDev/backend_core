package adapter

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
)

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
			return "", fmt.Errorf("%w: %w", port.ErrFailedToGenerateOtp, err)
		}
		code[i] = digits[n.Int64()]
	}
	return string(code), nil
}
