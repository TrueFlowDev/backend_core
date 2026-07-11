package port

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/valueobject"
)

type SmsOtpSender interface {
	Send(ctx context.Context, phone valueobject.Phone, otp entity.OTP) error
}
