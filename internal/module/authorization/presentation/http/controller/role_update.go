package controller

import (
	"errors"
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateRoleControllerInput struct {
	Title       string   `json:"title" binding:"required,min=2,max=100"`
	Permissions []string `json:"permissions" binding:"required,min=1,dive,required"`
} //	@name	UpdateRoleInput

type UpdateRoleController struct {
	usecase                 *usecase.UpdateRoleUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewUpdateRoleController(
	usecase *usecase.UpdateRoleUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *UpdateRoleController {
	return &UpdateRoleController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterUpdateRoleController(
	r *gin.Engine,
	controller *UpdateRoleController,
) {
	r.PATCH("/organization/:organization_id/role/:role_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("role.update"),
		controller.UpdateRole,
	)
}

// UpdateRole godoc
//
//	@Summary	Update role
//	@Tags		Role
//	@Accept		json
//	@Produce	json
//	@Param		organization_id	path	string						true	"Organization ID"
//	@Param		role_id			path	string						true	"Role ID"
//	@Param		request			body	UpdateRoleControllerInput	true	"Update role request"
//	@Success	204
//	@Failure	400	{object}	xerr.SwaggerErrOutput
//	@Failure	401	{object}	xerr.SwaggerErrOutput
//	@Failure	403	{object}	xerr.SwaggerErrOutput
//	@Failure	404	{object}	xerr.SwaggerErrOutput
//	@Failure	500	{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/role/{role_id} [patch]
func (c *UpdateRoleController) UpdateRole(ctx *gin.Context) {
	var input UpdateRoleControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("update role", validation.JSON))
		return
	}

	if err := c.usecase.Execute(ctx.Request.Context(), usecase.UpdateRoleInput{
		ID:             ctx.Param("role_id"),
		OrganizationID: ctx.Param("organization_id"),
		Title:          input.Title,
		Permissions:    input.Permissions,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
