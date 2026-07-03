package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/notification/domain/port"
)

type SendSmsInput struct {
	Phone   string
	Message string
}

type SendSMSUsecase struct {
	smsGateway port.SMSGateway
}

func NewSendSMSUsecase(smsGateway port.SMSGateway) *SendSMSUsecase {
	return &SendSMSUsecase{smsGateway: smsGateway}
}

func (u *SendSMSUsecase) Execute(ctx context.Context, input SendSmsInput) error {
	return u.smsGateway.Send(ctx, input.Phone, input.Message)
}
