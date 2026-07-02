package adapter

import (
	"context"
	"fmt"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
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

	return a.sendSMSUsecase.Execute(ctx, notificationUsecase.SendSmsInput{
		Phone:   phone.Value(),
		Message: message,
	})
}
