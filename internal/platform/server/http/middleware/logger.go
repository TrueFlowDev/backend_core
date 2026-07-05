package middleware

import (
	"net/http"
	"time"

	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/gin-gonic/gin"
)

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
