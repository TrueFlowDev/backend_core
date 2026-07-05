package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"
const RequestIDKey = "request_id"

type RequestID struct{}

func NewRequestID() *RequestID {
	return &RequestID{}
}

func (m *RequestID) Handle() gin.HandlerFunc {
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
