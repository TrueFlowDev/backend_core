package controller

import (
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/presentation/http/middleware"
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"github.com/gin-gonic/gin"
)

type SaveMyProfileControllerInput struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Headline  string `json:"headline,omitempty"`
	Bio       string `json:"bio,omitempty"`
} //	@name	SaveMyProfileInput

type SaveMyProfileController struct {
	usecase                 *usecase.SaveProfileUsecase
	authenticatedMiddleware *middleware.Authenticated
}

func NewSaveMyProfileController(
	usecase *usecase.SaveProfileUsecase,
	authenticatedMiddleware *middleware.Authenticated,
) *SaveMyProfileController {
	return &SaveMyProfileController{
		usecase:                 usecase,
		authenticatedMiddleware: authenticatedMiddleware,
	}
}

func RegisterSaveMyProfileController(
	r *gin.Engine,
	controller *SaveMyProfileController,
) {
	r.PUT(
		"/user/me/profile",
		controller.authenticatedMiddleware.Handle(),
		controller.SaveMyProfile,
	)
}

// SaveMyProfile godoc
//
//	@Summary		Create or update current user's profile
//	@Description	Creates or updates the authenticated user's profile.
//	@Tags			User
//	@Accept			json
//	@Param			request	body	SaveMyProfileControllerInput	true	"Profile"
//	@Success		204		"Profile updated successfully"
//	@Failure		400		{object}	xerr.SwaggerErrOutput
//	@Failure		401		{object}	xerr.SwaggerErrOutput
//	@Failure		500		{object}	xerr.SwaggerErrOutput
//	@Security		BearerAuth
//	@Router			/user/me/profile [put]
func (c *SaveMyProfileController) SaveMyProfile(ctx *gin.Context) {
	userID := ctx.MustGet(middleware.UserIDContextKey).(string)

	var input SaveMyProfileControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		_ = ctx.Error(err)
		return
	}

	_, err := c.usecase.Execute(
		ctx.Request.Context(),
		usecase.SaveProfileInput{
			UserID:    userID,
			Email:     input.Email,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Headline:  input.Headline,
			Bio:       input.Bio,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
