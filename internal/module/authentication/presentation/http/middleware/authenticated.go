package middleware

import (
	"strings"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/gin-gonic/gin"
)

const UserIDContextKey = "user_id"

type Authenticated struct {
	accessTokenProvider port.AccessTokenProvider
}

func NewAuthenticated(accessTokenProvider port.AccessTokenProvider) *Authenticated {
	return &Authenticated{
		accessTokenProvider: accessTokenProvider,
	}
}

func (m *Authenticated) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			_ = c.Error(port.ErrInvalidToken)
			c.Abort()
			return
		}

		const prefix = "Bearer "

		if !strings.HasPrefix(authHeader, prefix) {
			_ = c.Error(port.ErrInvalidToken)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, prefix)

		claims, err := m.accessTokenProvider.Verify(token)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		c.Set(UserIDContextKey, claims.UserID())

		c.Next()
	}
}
