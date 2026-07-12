package controller

import (
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/gin-gonic/gin"
)

type DashboardControllerOutput struct {
	EmployeeID       string `json:"employee_id"`
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	RoleID           string `json:"role_id"`
	JobTitle         string `json:"job_title"`
	EmploymentType   string `json:"employment_type"`
} //	@name	DashboardOutput

type ListMyDashboardsController struct {
	usecase                 *usecase.ListMyDashboardsUsecase
	authenticatedMiddleware *middleware.Authenticated
}

func NewListMyDashboardsController(
	usecase *usecase.ListMyDashboardsUsecase,
	authenticatedMiddleware *middleware.Authenticated,
) *ListMyDashboardsController {
	return &ListMyDashboardsController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterListMyDashboardsController(
	r *gin.Engine,
	controller *ListMyDashboardsController,
) {
	r.GET("/employee/me/dashboards", controller.authenticatedMiddleware.Handle(), controller.ListMyDashboards)
}

// ListMyDashboards godoc
//
//	@Summary		List my dashboards
//	@Description	Returns every active organization membership for the authenticated user.
//	@Tags			Employee
//	@Produce		json
//	@Success		200	{array}		DashboardControllerOutput
//	@Failure		401	{object}	xerr.SwaggerErrOutput
//	@Failure		500	{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/employee/me/dashboards [get]
func (c *ListMyDashboardsController) ListMyDashboards(ctx *gin.Context) {
	userID := ctx.MustGet(middleware.UserIDContextKey).(string)

	output, err := c.usecase.Execute(ctx.Request.Context(), usecase.ListMyDashboardsInput{
		UserID: userID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	result := make([]DashboardControllerOutput, len(output))
	for i, d := range output {
		result[i] = DashboardControllerOutput{
			EmployeeID:       d.EmployeeID,
			OrganizationID:   d.OrganizationID,
			OrganizationName: d.OrganizationName,
			RoleID:           d.RoleID,
			JobTitle:         d.JobTitle,
			EmploymentType:   d.EmploymentType,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
