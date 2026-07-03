package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

const RequestIDHeader = "X-Request-ID"
const RequestIDKey = "request_id"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set(RequestIDKey, reqID)
		c.Header(RequestIDHeader, reqID)

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if xe, ok := errors.AsType[*xerr.Error](err); ok {
			c.JSON(xe.HTTPStatus(), xe)
			return
		}

		c.JSON(http.StatusInternalServerError, xerr.New(xerr.CodeInternalError, xerr.WithMessage("unknown error")))
	}
}

func Logger(logger port.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		reqID, _ := c.Get(RequestIDKey)

		args := []any{
			"request_id", toString(reqID),
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"status", c.Writer.Status(),
			"latency", time.Since(start),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		}

		if len(c.Errors) > 0 {
			errorsLog := make([]string, 0, len(c.Errors))

			for _, e := range c.Errors {
				errorsLog = append(errorsLog, e.Err.Error())
			}

			args = append(args, "errors", errorsLog)
		}

		switch status := c.Writer.Status(); {
		case status >= http.StatusInternalServerError:
			logger.Error("http request", args...)
		case status >= http.StatusBadRequest:
			logger.Warn("http request", args...)
		default:
			logger.Info("http request", args...)
		}
	}
}

func toString(v any) string {
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

func NewGinEngine(cfg *config.Config, logger port.Logger) *gin.Engine {
	if cfg.App.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(
		gin.Recovery(),
		RequestID(),
		Logger(logger),
		ErrorHandler(),
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
						"error", err,
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
					"error", err,
				)
				return srv.Close()
			}

			logger.Info("http server stopped gracefully")
			return nil
		},
	})
}
