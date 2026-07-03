package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

var (
	ErrSmsOtpSendFailed = xerr.New(
		xerr.CodeInternalError,
		xerr.WithMessage("failed to send sms otp"),
	)
)

type SmsOtpSender interface {
	Send(ctx context.Context, phone value_object.Phone, otp entity.OTP) error
}
