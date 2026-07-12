package controller

import (
	"net/http"
	"time"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/organization/application/usecase"
	"github.com/gin-gonic/gin"
)

type FindOrganizationByIDControllerOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} //	@name	OrganizationOutput

type FindOrganizationByIDController struct {
	usecase                 *usecase.FindOrganizationByIDUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewFindOrganizationByIDController(
	usecase *usecase.FindOrganizationByIDUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *FindOrganizationByIDController {
	return &FindOrganizationByIDController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterFindOrganizationByIDController(
	r *gin.Engine,
	controller *FindOrganizationByIDController,
) {
	r.GET("/organization/:organization_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("organization.read"),
		controller.FindOrganizationByID,
	)
}

// FindOrganizationByID godoc
//
//	@Summary		Get organization
//	@Description	Returns the business details of an organization.
//	@Tags			Organization
//	@Produce		json
//	@Param			organization_id	path		string	true	"Organization ID"
//	@Success		200				{object}	FindOrganizationByIDControllerOutput
//	@Failure		401				{object}	xerr.SwaggerErrOutput
//	@Failure		403				{object}	xerr.SwaggerErrOutput
//	@Failure		404				{object}	xerr.SwaggerErrOutput
//	@Failure		500				{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/organization/{organization_id} [get]
func (c *FindOrganizationByIDController) FindOrganizationByID(ctx *gin.Context) {
	organization, err := c.usecase.Execute(ctx.Request.Context(), usecase.FindOrganizationByIDInput{
		ID: ctx.Param("organization_id"),
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, FindOrganizationByIDControllerOutput{
		ID:        organization.ID().Value(),
		Name:      organization.Name(),
		Category:  organization.Category().Value(),
		Active:    organization.Active(),
		CreatedAt: organization.CreatedAt(),
		UpdatedAt: organization.UpdatedAt(),
	})
}
