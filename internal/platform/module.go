package platform

import (
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		config.NewConfig,
	),
	fx.Invoke(
		config.LoadFromEnvFile,
	),
)
