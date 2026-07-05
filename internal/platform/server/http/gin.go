package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinEngine(cfg *config.Config, log port.Logger) *gin.Engine {
	if cfg.App.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(
		gin.Recovery(),
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.ErrorHandler(),
	)

	return router
}

func NewHTTPServer(
	cfg *config.Config,
	engine *gin.Engine,
) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.Port),
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func StartHTTPServer(lc fx.Lifecycle, srv *http.Server, logger port.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return fmt.Errorf("failed to listen on %s: %w", srv.Addr, err)
			}

			go func() {
				logger.Info(
					"http server started",
					"addr", srv.Addr,
				)
				if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error(
						"http server crashed",
						"error_handling", err,
					)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down http server")

			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(shutdownCtx); err != nil {
				logger.Error(
					"graceful shutdown failed, forcing close",
					"error_handling", err,
				)
				return srv.Close()
			}

			logger.Info("http server stopped gracefully")
			return nil
		},
	})
}
