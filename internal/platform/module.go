package platform

import (
	"github.com/TrueFlowDev/Backend/internal/platform/cache"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/platform/database"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		config.NewConfig,
		database.NewPostgres,
		cache.NewRedis,
	),
	fx.Invoke(
		config.LoadFromEnvFile,
		http.RegisterSwagger,
	),
)
