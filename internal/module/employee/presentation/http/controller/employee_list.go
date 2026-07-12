package controller

import (
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/gin-gonic/gin"
)

type EmployeeControllerOutput struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	RoleID           string `json:"role_id"`
	JobTitle         string `json:"job_title"`
	MembershipStatus string `json:"membership_status"`
	EmploymentType   string `json:"employment_type"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
} //	@name	EmployeeOutput

type ListEmployeesController struct {
	usecase                 *usecase.ListEmployeesUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewListEmployeesController(
	usecase *usecase.ListEmployeesUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *ListEmployeesController {
	return &ListEmployeesController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterListEmployeesController(r *gin.Engine, controller *ListEmployeesController) {
	r.GET("/organization/:organization_id/employee",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("employee.list"),
		controller.ListEmployees,
	)
}

// ListEmployees godoc
//
//	@Summary	List employees
//	@Tags		Employee
//	@Produce	json
//	@Param		organization_id	path		string	true	"Organization ID"
//	@Success	200				{array}		EmployeeControllerOutput
//	@Failure	401				{object}	xerr.SwaggerErrOutput
//	@Failure	403				{object}	xerr.SwaggerErrOutput
//	@Failure	500				{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/employee [get]
func (c *ListEmployeesController) ListEmployees(ctx *gin.Context) {
	employees, err := c.usecase.Execute(ctx.Request.Context(), usecase.ListEmployeesInput{
		OrganizationID: ctx.Param("organization_id"),
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	result := make([]EmployeeControllerOutput, len(employees))
	for i, e := range employees {
		result[i] = EmployeeControllerOutput{
			ID:               e.ID,
			UserID:           e.UserID,
			RoleID:           e.RoleID,
			JobTitle:         e.JobTitle,
			MembershipStatus: e.MembershipStatus,
			EmploymentType:   e.EmploymentType,
			FirstName:        e.FirstName,
			LastName:         e.LastName,
			Phone:            e.Phone,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
