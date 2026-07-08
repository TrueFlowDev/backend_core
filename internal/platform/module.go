package platform

import (
	"github.com/TrueFlowDev/Backend/internal/platform/cache"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/platform/database"
	"github.com/TrueFlowDev/Backend/internal/platform/logger"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		fx.Annotate(
			logger.NewZapLogger,
			fx.As(new(port.Logger)),
		),
		fx.Annotate(
			database.NewTxManager,
			fx.As(new(port.TxManager)),
		),
		config.NewConfig,
		database.NewPostgres,
		cache.NewRedis,
		middleware.NewErrorHandler,
		middleware.NewLogger,
		middleware.NewRequestID,
	),
	fx.Invoke(
		config.LoadFromEnvFile,
		http.RegisterSwagger,
	),
)
