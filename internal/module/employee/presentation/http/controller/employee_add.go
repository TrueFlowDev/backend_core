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

type AddEmployeeControllerInput struct {
	Phone            string `json:"phone" binding:"required"`
	RoleID           string `json:"role_id" binding:"required"`
	JobTitle         string `json:"job_title" binding:"required"`
	MembershipStatus string `json:"membership_status" binding:"required"`
	EmploymentType   string `json:"employment_type" binding:"required"`
} //	@name	AddEmployeeInput

type AddEmployeeControllerOutput struct {
	ID string `json:"id"`
} //	@name	AddEmployeeOutput

type AddEmployeeController struct {
	usecase                 *usecase.AddEmployeeUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewAddEmployeeController(
	usecase *usecase.AddEmployeeUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *AddEmployeeController {
	return &AddEmployeeController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterAddEmployeeController(r *gin.Engine, controller *AddEmployeeController) {
	r.POST("/organization/:organization_id/employee",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("employee.create"),
		controller.AddEmployee,
	)
}

// AddEmployee godoc
//
//	@Summary	Add employee
//	@Tags		Employee
//	@Accept		json
//	@Produce	json
//	@Param		organization_id	path		string						true	"Organization ID"
//	@Param		request			body		AddEmployeeControllerInput	true	"Add employee request"
//	@Success	201				{object}	AddEmployeeControllerOutput
//	@Failure	400				{object}	xerr.SwaggerErrOutput
//	@Failure	401				{object}	xerr.SwaggerErrOutput
//	@Failure	403				{object}	xerr.SwaggerErrOutput
//	@Failure	404				{object}	xerr.SwaggerErrOutput
//	@Failure	409				{object}	xerr.SwaggerErrOutput
//	@Failure	500				{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/employee [post]
func (c *AddEmployeeController) AddEmployee(ctx *gin.Context) {
	var input AddEmployeeControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError("add employee", validation.JSON))
		return
	}

	output, err := c.usecase.Execute(ctx.Request.Context(), usecase.AddEmployeeInput{
		OrganizationID:   ctx.Param("organization_id"),
		Phone:            input.Phone,
		RoleID:           input.RoleID,
		JobTitle:         input.JobTitle,
		MembershipStatus: input.MembershipStatus,
		EmploymentType:   input.EmploymentType,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, AddEmployeeControllerOutput{ID: output.ID})
}
