package platform

import (
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/platform/database"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		config.NewConfig,
		database.NewPostgres,
	),
	fx.Invoke(
		config.LoadFromEnvFile,
		database.ConnectPostgres,
	),
)
