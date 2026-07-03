package notification

import (
	"github.com/TrueFlowDev/Backend/internal/module/notification/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/notification/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/notification/infrastructure/adapter"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"notification",
	fx.Provide(
		usecase.NewSendSMSUsecase,
		fx.Annotate(
			adapter.NewLocalSmsGateway,
			fx.As(new(port.SMSGateway)),
		),
	),
)
