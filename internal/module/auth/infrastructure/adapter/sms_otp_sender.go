package adapter

import (
	"context"
	"fmt"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
	notificationUsecase "github.com/TrueFlowDev/Backend/internal/module/notification/application/usecase"
)

type SmsOtpSenderAdapter struct {
	sendSMSUsecase *notificationUsecase.SendSMSUsecase
}

func NewSmsOtpSenderAdapter(sendSMSUsecase *notificationUsecase.SendSMSUsecase) *SmsOtpSenderAdapter {
	return &SmsOtpSenderAdapter{sendSMSUsecase: sendSMSUsecase}
}

func (a *SmsOtpSenderAdapter) Send(ctx context.Context, phone value_object.Phone, otp entity.OTP) error {
	message := fmt.Sprintf("OTP Code: %s", otp.Code())

	if err := a.sendSMSUsecase.Execute(ctx, notificationUsecase.SendSmsInput{
		Phone:   phone.Value(),
		Message: message,
	}); err != nil {
		return xerr.Wrap(err, port.ErrSmsOtpSendFailed.Code())
	}

	return nil
}
