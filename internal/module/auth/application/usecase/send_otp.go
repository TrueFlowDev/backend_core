package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"go.uber.org/fx"
)

type SendOtpInput struct {
	Phone string
}

type SendOtpUsecase struct {
	otpCodeGenerator  port.OtpCodeGenerator
	otpStore          port.OTPStore
	otpSender         port.SmsOtpSender
	otpTTL            time.Duration
	userFinderByPhone port.UserFinderByPhone
}

type SendOtpParams struct {
	fx.In

	Config            *config.Config
	OtpCodeGenerator  port.OtpCodeGenerator
	OtpStore          port.OTPStore
	OtpSender         port.SmsOtpSender
	UserFinderByPhone port.UserFinderByPhone
}

func NewSendOtpUsecase(p SendOtpParams) *SendOtpUsecase {
	return &SendOtpUsecase{
		otpCodeGenerator:  p.OtpCodeGenerator,
		otpStore:          p.OtpStore,
		otpSender:         p.OtpSender,
		otpTTL:            p.Config.OTP.TTL,
		userFinderByPhone: p.UserFinderByPhone,
	}
}

func (u *SendOtpUsecase) Execute(ctx context.Context, input SendOtpInput) error {
	_, err := u.userFinderByPhone.FindByPhone(
		ctx,
		port.UserFinderByPhoneInput{Phone: input.Phone},
	)
	if err != nil {
		if !errors.Is(err, port.ErrUserNotFound) {
			return err
		}
	} else {
		return port.ErrUserAlreadyExists
	}

	phone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return err
	}

	otpCode, err := u.otpCodeGenerator.Generate()
	if err != nil {
		return err
	}

	otpExpiresAt := time.Now().UTC().Add(u.otpTTL)
	otp, err := entity.NewOTP(otpCode, otpExpiresAt)
	if err != nil {
		return err
	}

	if err := u.otpStore.Set(ctx, phone, otp); err != nil {
		return err
	}

	if err := u.otpSender.Send(ctx, phone, otp); err != nil {
		return err
	}

	return nil
}
