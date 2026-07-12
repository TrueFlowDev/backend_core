package controller

import (
	"net/http"

	authMiddleware "github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	authzMiddleware "github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/gin-gonic/gin"
)

type RemoveEmployeeController struct {
	usecase                 *usecase.RemoveEmployeeUsecase
	authenticatedMiddleware *authMiddleware.Authenticated
	permissionGuard         *authzMiddleware.PermissionGuard
}

func NewRemoveEmployeeController(
	usecase *usecase.RemoveEmployeeUsecase,
	authenticatedMiddleware *authMiddleware.Authenticated,
	permissionGuard *authzMiddleware.PermissionGuard,
) *RemoveEmployeeController {
	return &RemoveEmployeeController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
		permissionGuard:         permissionGuard,
	}
}

func RegisterRemoveEmployeeController(r *gin.Engine, controller *RemoveEmployeeController) {
	r.DELETE("/organization/:organization_id/employee/:employee_id",
		controller.authenticatedMiddleware.Handle(),
		controller.permissionGuard.Require("employee.delete"),
		controller.RemoveEmployee,
	)
}

// RemoveEmployee godoc
//
//	@Summary	Remove employee
//	@Tags		Employee
//	@Param		organization_id	path	string	true	"Organization ID"
//	@Param		employee_id		path	string	true	"Employee ID"
//	@Success	204
//	@Failure	400	{object}	xerr.SwaggerErrOutput
//	@Failure	401	{object}	xerr.SwaggerErrOutput
//	@Failure	403	{object}	xerr.SwaggerErrOutput
//	@Failure	404	{object}	xerr.SwaggerErrOutput
//	@Failure	500	{object}	xerr.SwaggerErrOutput
//	@Security	BearerAuth
//	@Router		/organization/{organization_id}/employee/{employee_id} [delete]
func (c *RemoveEmployeeController) RemoveEmployee(ctx *gin.Context) {
	userID := ctx.MustGet(authMiddleware.UserIDContextKey).(string)

	if err := c.usecase.Execute(ctx.Request.Context(), usecase.RemoveEmployeeInput{
		ID:               ctx.Param("employee_id"),
		OrganizationID:   ctx.Param("organization_id"),
		RequestingUserID: userID,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
