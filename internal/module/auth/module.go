package auth

import (
	"github.com/TrueFlowDev/Backend/internal/module/auth/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/auth/presentation/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth",
	fx.Provide(
		fx.Annotate(
			adapter.NewSmsOtpSenderAdapter,
			fx.As(new(port.SmsOtpSender)),
		),
		fx.Annotate(
			adapter.NewPasswordHasher,
			fx.As(new(port.PasswordHasher)),
		),
		fx.Annotate(
			adapter.NewOtpCodeGenerator,
			fx.As(new(port.OtpCodeGenerator)),
		),
		fx.Annotate(
			adapter.NewJwtProvider,
			fx.As(new(port.AccessTokenProvider)),
		),
		fx.Annotate(
			adapter.NewOTPStore,
			fx.As(new(port.OTPStore)),
		),
		usecase.NewSendOtpUsecase,
		usecase.NewVerifyOTPAndRegisterUsecase,
		http.NewSendOtpController,
		http.NewVerifyOTPAndRegisterController,
	),
	fx.Invoke(
		http.RegisterSendOtpController,
		http.RegisterVerifyOTPAndRegisterController,
	),
)
