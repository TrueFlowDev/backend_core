package port

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/value_object"
)

type SmsOtpSender interface {
	Send(ctx context.Context, phone value_object.Phone, otp entity.OTP) error
}
