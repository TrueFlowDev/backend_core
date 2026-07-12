package controller

import (
	"errors"
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/organization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateOrganizationControllerInput struct {
	Category            string `json:"category" binding:"required,oneof=technology finance retail manufacturing other"`
	Name                string `json:"name" binding:"required,min=3,max=100"`
	OwnerJobTitle       string `json:"owner_job_title" binding:"required,min=2,max=100"`
	OwnerEmploymentType string `json:"owner_employment_type" binding:"required,oneof=full_time part_time contract"`
} //	@name	CreateOrganizationInput

type CreateOrganizationControllerOutput struct {
	OrganizationID string `json:"organization_id"`
	RoleID         string `json:"role_id"`
	EmployeeID     string `json:"employee_id"`
} //	@name	CreateOrganizationOutput

type CreateOrganizationController struct {
	usecase                 *usecase.CreateOrganizationWithOwnerUsecase
	authenticatedMiddleware *middleware.Authenticated
}

func NewCreateOrganizationController(
	usecase *usecase.CreateOrganizationWithOwnerUsecase,
	authenticatedMiddleware *middleware.Authenticated,
) *CreateOrganizationController {
	return &CreateOrganizationController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterCreateOrganizationController(
	r *gin.Engine,
	controller *CreateOrganizationController,
) {
	r.POST("/organization", controller.authenticatedMiddleware.Handle(), controller.CreateOrganization)
}

// CreateOrganization godoc
//
//	@Summary		Create organization
//	@Description	Creates a new organization and makes the caller its owner.
//	@Tags			Organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateOrganizationControllerInput	true	"Create organization request"
//	@Success		201		{object}	CreateOrganizationControllerOutput
//	@Failure		400		{object}	xerr.SwaggerErrOutput
//	@Failure		401		{object}	xerr.SwaggerErrOutput
//	@Failure		500		{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/organization [post]
func (c *CreateOrganizationController) CreateOrganization(ctx *gin.Context) {
	var input CreateOrganizationControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("create organization", validation.JSON))
		return
	}

	userID := ctx.MustGet(middleware.UserIDContextKey).(string)

	output, err := c.usecase.Execute(ctx.Request.Context(), usecase.CreateOrganizationWithOwnerInput{
		Category:            input.Category,
		Name:                input.Name,
		OwnerUserID:         userID,
		OwnerJobTitle:       input.OwnerJobTitle,
		OwnerEmploymentType: input.OwnerEmploymentType,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, CreateOrganizationControllerOutput{
		OrganizationID: output.OrganizationID,
		RoleID:         output.RoleID,
		EmployeeID:     output.EmployeeID,
	})
}
