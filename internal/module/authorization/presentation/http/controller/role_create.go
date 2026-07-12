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

type CreateRoleControllerInput struct {
	Title       string   `json:"title" binding:"required,min=2,max=100"`
	Permissions []string `json:"permissions" binding:"required,min=1,dive,required"`
} //	@name	CreateRoleInput

type CreateRoleControllerOutput struct {
	ID string `json:"id"`
} //	@name	CreateRoleOutput

type CreateRoleController struct {
	usecase                 *usecase.CreateRoleUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewCreateRoleController(
	usecase *usecase.CreateRoleUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *CreateRoleController {
	return &CreateRoleController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterCreateRoleController(
	r *gin.Engine,
	controller *CreateRoleController,
) {
	r.POST("/organization/:organization_id/role",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("role.create"),
		controller.CreateRole,
	)
}

// CreateRole godoc
//
//	@Summary		Create role
//	@Description	Creates a new role for an organization.
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			organization_id	path		string						true	"Organization ID"
//	@Param			request			body		CreateRoleControllerInput	true	"Create role request"
//	@Success		201				{object}	CreateRoleControllerOutput
//	@Failure		400				{object}	xerr.SwaggerErrOutput
//	@Failure		401				{object}	xerr.SwaggerErrOutput
//	@Failure		403				{object}	xerr.SwaggerErrOutput
//	@Failure		500				{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/organization/{organization_id}/role [post]
func (c *CreateRoleController) CreateRole(ctx *gin.Context) {
	var input CreateRoleControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("create role", validation.JSON))
		return
	}

	organizationID := ctx.Param("organization_id")

	output, err := c.usecase.Execute(ctx.Request.Context(), usecase.CreateRoleInput{
		OrganizationID: organizationID,
		Title:          input.Title,
		Permissions:    input.Permissions,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, CreateRoleControllerOutput{ID: output.ID})
}
