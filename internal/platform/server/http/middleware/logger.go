package middleware

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/gin-gonic/gin"
)

func Logger(logger port.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/healthz" {
			c.Next()
			return
		}

		start := time.Now()

		c.Next()

		reqID, exists := c.Get(RequestIDKey)
		reqIDStr := "unknown"
		if exists {
			reqIDStr = toString(reqID)
		}

		args := []any{
			"request_id", reqIDStr,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", sanitizeQuery(c.Request.URL.RawQuery),
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

var sensitiveQueryParams = map[string]bool{
	"token": true, "code": true, "password": true, "secret": true, "otp": true,
}

func sanitizeQuery(rawQuery string) string {
	values, err := url.ParseQuery(rawQuery)
	if err != nil {
		return ""
	}
	for key := range values {
		if sensitiveQueryParams[strings.ToLower(key)] {
			values.Set(key, "***")
		}
	}
	return values.Encode()
}
