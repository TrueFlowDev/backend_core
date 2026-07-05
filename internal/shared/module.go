package shared

import (
	"github.com/TrueFlowDev/Backend/internal/platform/server/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	fx.Provide(
		http.NewGinEngine,
		http.NewHTTPServer,
	),
	fx.Invoke(
		http.StartHTTPServer,
	),
)
