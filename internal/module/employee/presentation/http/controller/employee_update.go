package controller

import (
	"errors"
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateEmployeeControllerInput struct {
	RoleID           string `json:"role_id" binding:"required"`
	JobTitle         string `json:"job_title" binding:"required"`
	MembershipStatus string `json:"membership_status" binding:"required"`
	EmploymentType   string `json:"employment_type" binding:"required"`
} //	@name	UpdateEmployeeInput

type UpdateEmployeeController struct {
	usecase                 *usecase.UpdateEmployeeUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewUpdateEmployeeController(
	usecase *usecase.UpdateEmployeeUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *UpdateEmployeeController {
	return &UpdateEmployeeController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterUpdateEmployeeController(r *gin.Engine, controller *UpdateEmployeeController) {
	r.PATCH("/organization/:organization_id/employee/:employee_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("employee.update"),
		controller.UpdateEmployee,
	)
}

// UpdateEmployee godoc
//
//	@Summary	Update employee
//	@Tags		Employee
//	@Accept		json
//	@Produce	json
//	@Param		organization_id	path	string							true	"Organization ID"
//	@Param		employee_id		path	string							true	"Employee ID"
//	@Param		request			body	UpdateEmployeeControllerInput	true	"Update employee request"
//	@Success	204
//	@Failure	400	{object}	xerr.SwaggerErrOutput
//	@Failure	401	{object}	xerr.SwaggerErrOutput
//	@Failure	403	{object}	xerr.SwaggerErrOutput
//	@Failure	404	{object}	xerr.SwaggerErrOutput
//	@Failure	500	{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/employee/{employee_id} [patch]
func (c *UpdateEmployeeController) UpdateEmployee(ctx *gin.Context) {
	userID := ctx.MustGet(authMiddleware.UserIDContextKey).(string)

	var input UpdateEmployeeControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("update employee", validation.JSON))
		return
	}

	if err := c.usecase.Execute(ctx.Request.Context(), usecase.UpdateEmployeeInput{
		ID:               ctx.Param("employee_id"),
		OrganizationID:   ctx.Param("organization_id"),
		RoleID:           input.RoleID,
		JobTitle:         input.JobTitle,
		MembershipStatus: input.MembershipStatus,
		EmploymentType:   input.EmploymentType,
		RequestingUserID: userID,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
