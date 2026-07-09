package controller

import (
	"net/http"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"github.com/gin-gonic/gin"
)

type GetMyProfileControllerOutput struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Headline  string `json:"headline,omitempty"`
	Bio       string `json:"bio,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} //	@name	GetMyProfileOutput

type GetMyProfileController struct {
	usecase                 *usecase.FindProfileByUserIDUsecase
	authenticatedMiddleware *middleware.Authenticated
}

func NewGetMyProfileController(
	usecase *usecase.FindProfileByUserIDUsecase,
	authenticatedMiddleware *middleware.Authenticated,
) *GetMyProfileController {
	return &GetMyProfileController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterGetMyProfileController(
	r *gin.Engine,
	controller *GetMyProfileController,
) {
	r.GET(
		"/user/me/profile",
		controller.authenticatedMiddleware.Handle(),
		controller.GetMyProfile,
	)
}

// GetMyProfile godoc
//
//	@Summary		Get current user's profile
//	@Description	Returns the authenticated user's profile information.
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	GetMyProfileControllerOutput
//	@Failure		401	{object}	xerr.SwaggerErrOutput
//	@Failure		404	{object}	xerr.SwaggerErrOutput
//	@Failure		500	{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/user/me/profile [get]
func (c *GetMyProfileController) GetMyProfile(ctx *gin.Context) {
	userID := ctx.MustGet(middleware.UserIDContextKey).(string)

	profile, err := c.usecase.Execute(
		ctx.Request.Context(),
		usecase.FindProfileByUserIDInput{
			UserID: userID,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, GetMyProfileControllerOutput{
		UserID:    profile.UserID().Value(),
		Email:     profile.Email().Value(),
		FirstName: profile.FirstName(),
		LastName:  profile.LastName(),
		Headline:  profile.Headline(),
		Bio:       profile.Bio(),
		CreatedAt: profile.CreatedAt(),
		UpdatedAt: profile.UpdatedAt(),
	})
}
