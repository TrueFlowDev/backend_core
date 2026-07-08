package shared

import (
	"github.com/TrueFlowDev/Backend/internal/platform/server/http"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	fx.Provide(
		http.NewGinEngine,
		http.NewHTTPServer,
		database.NewBaseRepo,
	),
	fx.Invoke(
		http.StartHTTPServer,
	),
)
