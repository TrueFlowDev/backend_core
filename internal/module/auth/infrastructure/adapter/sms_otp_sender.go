package adapter

import (
	"context"
	"fmt"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
	notificationUsecase "github.com/TrueFlowDev/Backend/internal/module/notification/application/usecase"
)

var _ port.SmsOtpSender = (*SmsOtpSender)(nil)

type SmsOtpSender struct {
	sendSMSUsecase *notificationUsecase.SendSMSUsecase
}

func NewSmsOtpSender(sendSMSUsecase *notificationUsecase.SendSMSUsecase) *SmsOtpSender {
	return &SmsOtpSender{sendSMSUsecase: sendSMSUsecase}
}

func (a *SmsOtpSender) Send(ctx context.Context, phone value_object.Phone, otp entity.OTP) error {
	message := fmt.Sprintf("OTP Code: %s", otp.Code())

	if err := a.sendSMSUsecase.Execute(ctx, notificationUsecase.SendSmsInput{
		Phone:   phone.Value(),
		Message: message,
	}); err != nil {
		return err
	}

	return nil
}
