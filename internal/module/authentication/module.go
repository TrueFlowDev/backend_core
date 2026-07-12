package authentication

import (
	"github.com/TrueFlowDev/Backend/internal/module/authentication/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/controller"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"authentication",
	fx.Provide(
		fx.Annotate(
			adapter.NewSmsOtpSender,
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
		fx.Annotate(
			adapter.NewUserRegisterer,
			fx.As(new(port.UserRegisterer)),
		),
		fx.Annotate(
			adapter.NewUserFinderByPhone,
			fx.As(new(port.UserFinderByPhone)),
		),
		usecase.NewSendOtpUsecase,
		usecase.NewVerifyOTPAndRegisterUsecase,
		usecase.NewLoginUsecase,
		controller.NewSendOtpController,
		controller.NewVerifyOTPAndRegisterController,
		controller.NewLoginController,
		middleware.NewAuthenticated,
	),
	fx.Invoke(
		controller.RegisterSendOtpController,
		controller.RegisterVerifyOTPAndRegisterController,
		controller.RegisterLoginController,
	),
)
