package controller

import (
	"errors"
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/auth/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginControllerInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
} //	@name	LoginInput

type LoginControllerOutput struct {
	AccessToken string `json:"access_token"`
} //	@name	LoginOutput

type LoginController struct {
	usecase *usecase.LoginUsecase
}

func NewLoginController(
	usecase *usecase.LoginUsecase,
) *LoginController {
	return &LoginController{
		usecase: usecase,
	}
}

func RegisterLoginController(
	r *gin.Engine,
	controller *LoginController,
) {
	r.POST("/auth/login", controller.Login)
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Logins the user and returns an access token.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginControllerInput	true	"Register request"
//	@Success		200		{object}	LoginControllerOutput
//	@Failure		400		{object}	xerr.SwaggerErrOutput
//	@Failure		500		{object}	xerr.SwaggerErrOutput
//	@Router			/auth/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	var input LoginControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError(
			"login",
			validation.JSON,
		))
		return
	}

	output, err := c.usecase.Execute(
		ctx.Request.Context(),
		usecase.LoginInput{
			Phone:    input.Phone,
			Password: input.Password,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, LoginControllerOutput{
		AccessToken: output.AccessToken,
	})
}
