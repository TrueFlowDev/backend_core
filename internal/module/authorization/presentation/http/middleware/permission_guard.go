package middleware

import (
	"github.com/Ali127Dev/xerr"
	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/gin-gonic/gin"
)

var ErrPermissionDenied = xerr.New(xerr.CodeForbidden)

type PermissionGuard struct {
	hasPermissionUsecase *usecase.HasPermissionUsecase
}

func NewPermissionGuard(hasPermissionUsecase *usecase.HasPermissionUsecase) *PermissionGuard {
	return &PermissionGuard{hasPermissionUsecase: hasPermissionUsecase}
}

func (g *PermissionGuard) Require(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet(authMiddleware.UserIDContextKey).(string)
		organizationID := c.Param("organization_id")

		allowed, err := g.hasPermissionUsecase.Execute(c.Request.Context(), usecase.HasPermissionInput{
			UserID:         userID,
			OrganizationID: organizationID,
			Permission:     permission,
		})
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		if !allowed {
			_ = c.Error(ErrPermissionDenied)
			c.Abort()
			return
		}

		c.Next()
	}
}
