package auth

import (
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/infrastructure/adapter"
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
	),
)
