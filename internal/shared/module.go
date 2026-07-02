package shared

import (
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/adapter"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	fx.Provide(
		fx.Annotate(
			adapter.NewZapLogger,
			fx.As(new(port.Logger)),
		),
	),
)
