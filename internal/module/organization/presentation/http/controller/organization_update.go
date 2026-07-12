package controller

import (
	"errors"
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/organization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateOrganizationControllerInput struct {
	Category string `json:"category" binding:"required,oneof=technology finance retail manufacturing other"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
} //	@name	UpdateOrganizationInput

type UpdateOrganizationController struct {
	usecase                 *usecase.UpdateOrganizationUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewUpdateOrganizationController(
	usecase *usecase.UpdateOrganizationUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *UpdateOrganizationController {
	return &UpdateOrganizationController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterUpdateOrganizationController(
	r *gin.Engine,
	controller *UpdateOrganizationController,
) {
	r.PATCH("/organization/:organization_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("organization.update"),
		controller.UpdateOrganization,
	)
}

// UpdateOrganization godoc
//
//	@Summary	Update organization
//	@Tags		Organization
//	@Accept		json
//	@Produce	json
//	@Param		organization_id	path	string								true	"Organization ID"
//	@Param		request			body	UpdateOrganizationControllerInput	true	"Update organization request"
//	@Success	204
//	@Failure	400	{object}	xerr.SwaggerErrOutput
//	@Failure	401	{object}	xerr.SwaggerErrOutput
//	@Failure	403	{object}	xerr.SwaggerErrOutput
//	@Failure	404	{object}	xerr.SwaggerErrOutput
//	@Failure	500	{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id} [patch]
func (c *UpdateOrganizationController) UpdateOrganization(ctx *gin.Context) {
	var input UpdateOrganizationControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("update organization", validation.JSON))
		return
	}

	if err := c.usecase.Execute(ctx.Request.Context(), usecase.UpdateOrganizationInput{
		ID:       ctx.Param("organization_id"),
		Name:     input.Name,
		Category: input.Category,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
