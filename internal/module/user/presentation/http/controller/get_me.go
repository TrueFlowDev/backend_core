package controller

import (
	"net/http"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"github.com/gin-gonic/gin"
)

type GetMeControllerOutput struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} //	@name	GetMeOutput

type GetMeController struct {
	usecase                 *usecase.FindUserByIDUsecase
	authenticatedMiddleware *middleware.Authenticated
}

func NewGetMeController(
	usecase *usecase.FindUserByIDUsecase,
	authenticatedMiddleware *middleware.Authenticated,
) *GetMeController {
	return &GetMeController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterGetMeController(
	r *gin.Engine,
	controller *GetMeController,
) {
	r.GET("/user/me", controller.authenticatedMiddleware.Handle(), controller.GetMe)
}

// GetMe godoc
//
//	@Summary		Get current user
//	@Description	Returns the authenticated user's profile.
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	GetMeControllerOutput
//	@Failure		401	{object}	xerr.SwaggerErrOutput
//	@Failure		500	{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/user/me [get]
func (c *GetMeController) GetMe(ctx *gin.Context) {
	userID := ctx.MustGet(middleware.UserIDContextKey).(string)

	output, err := c.usecase.Execute(
		ctx.Request.Context(),
		usecase.FindUserByIDInput{
			ID: userID,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, GetMeControllerOutput{
		ID:        output.ID().Value(),
		Phone:     output.Phone().Value(),
		CreatedAt: output.CreatedAt(),
		UpdatedAt: output.UpdatedAt(),
	})
}
