package controller

import (
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/gin-gonic/gin"
)

type DeleteRoleController struct {
	usecase                 *usecase.DeleteRoleUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewDeleteRoleController(
	usecase *usecase.DeleteRoleUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *DeleteRoleController {
	return &DeleteRoleController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterDeleteRoleController(
	r *gin.Engine,
	controller *DeleteRoleController,
) {
	r.DELETE("/organization/:organization_id/role/:role_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("role.delete"),
		controller.DeleteRole,
	)
}

// DeleteRole godoc
//
//	@Summary	Delete role
//	@Tags		Role
//	@Param		organization_id	path	string	true	"Organization ID"
//	@Param		role_id			path	string	true	"Role ID"
//	@Success	204
//	@Failure	400	{object}	xerr.SwaggerErrOutput
//	@Failure	401	{object}	xerr.SwaggerErrOutput
//	@Failure	403	{object}	xerr.SwaggerErrOutput
//	@Failure	404	{object}	xerr.SwaggerErrOutput
//	@Failure	500	{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/role/{role_id} [delete]
func (c *DeleteRoleController) DeleteRole(ctx *gin.Context) {
	if err := c.usecase.Execute(ctx.Request.Context(), usecase.DeleteRoleInput{
		ID:             ctx.Param("role_id"),
		OrganizationID: ctx.Param("organization_id"),
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
