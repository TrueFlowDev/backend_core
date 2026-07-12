package controller

import (
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/gin-gonic/gin"
)

type PermissionOutput struct {
	Value string `json:"value"`
	Title string `json:"title"`
} //	@name	PermissionOutput

type ListPermissionsControllerOutput struct {
	Category    string             `json:"category"`
	Permissions []PermissionOutput `json:"permissions"`
} //	@name	ListPermissionsOutput

type ListPermissionsController struct {
	usecase                 *usecase.ListPermissionsUseCase
	authenticatedMiddleware *middleware.Authenticated
}

func NewListPermissionsController(
	usecase *usecase.ListPermissionsUseCase,
	authenticatedMiddleware *middleware.Authenticated,
) *ListPermissionsController {
	return &ListPermissionsController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterListPermissionsController(
	r *gin.Engine,
	controller *ListPermissionsController,
) {
	r.GET("/permission", controller.authenticatedMiddleware.Handle(), controller.ListPermissions)
}

// ListPermissions godoc
//
//	@Summary		List permissions
//	@Description	Returns all permissions grouped by category.
//	@Tags			Authorization
//	@Produce		json
//	@Success		200	{array}		ListPermissionsControllerOutput
//	@Failure		401	{object}	xerr.SwaggerErrOutput
//	@Failure		500	{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/permission [get]
func (c *ListPermissionsController) ListPermissions(ctx *gin.Context) {
	output := c.usecase.Execute(ctx.Request.Context())

	result := make([]ListPermissionsControllerOutput, len(output))
	for i, group := range output {
		permissions := make([]PermissionOutput, len(group.Permissions))
		for j, p := range group.Permissions {
			permissions[j] = PermissionOutput{
				Value: p.Value,
				Title: p.Title,
			}
		}
		result[i] = ListPermissionsControllerOutput{
			Category:    group.Category,
			Permissions: permissions,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
