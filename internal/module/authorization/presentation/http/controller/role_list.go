package controller

import (
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/gin-gonic/gin"
)

type RoleControllerOutput struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	IsOwner     bool     `json:"is_owner"`
	Permissions []string `json:"permissions"`
} //	@name	RoleOutput

type ListRolesController struct {
	usecase                 *usecase.ListRolesUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewListRolesController(
	usecase *usecase.ListRolesUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *ListRolesController {
	return &ListRolesController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterListRolesController(
	r *gin.Engine,
	controller *ListRolesController,
) {
	r.GET("/organization/:organization_id/role",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("role.list"),
		controller.ListRoles,
	)
}

// ListRoles godoc
//
//	@Summary		List roles
//	@Description	Returns every role defined for an organization.
//	@Tags			Role
//	@Produce		json
//	@Param			organization_id	path		string	true	"Organization ID"
//	@Success		200				{array}		RoleControllerOutput
//	@Failure		401				{object}	xerr.SwaggerErrOutput
//	@Failure		403				{object}	xerr.SwaggerErrOutput
//	@Failure		500				{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/organization/{organization_id}/role [get]
func (c *ListRolesController) ListRoles(ctx *gin.Context) {
	organizationID := ctx.Param("organization_id")

	roles, err := c.usecase.Execute(ctx.Request.Context(), usecase.ListRolesInput{
		OrganizationID: organizationID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	result := make([]RoleControllerOutput, len(roles))
	for i, role := range roles {
		permissions := make([]string, len(role.Permissions()))
		for j, p := range role.Permissions() {
			permissions[j] = p.Value()
		}

		result[i] = RoleControllerOutput{
			ID:          role.ID().Value(),
			Title:       role.Title(),
			IsOwner:     role.IsOwner(),
			Permissions: permissions,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
