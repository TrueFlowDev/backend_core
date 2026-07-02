package platform

import (
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Invoke(
		config.LoadFromEnvFile,
	),
)
