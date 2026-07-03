package http

import (
	"errors"
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/auth/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/shared/presentation/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VerifyOTPAndRegisterControllerInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
} //	@name	VerifyOTPAndRegisterInput

type VerifyOTPAndRegisterControllerOutput struct {
	AccessToken string `json:"accessToken"`
} //	@name	VerifyOTPAndRegisterOutput

type VerifyOTPAndRegisterController struct {
	usecase *usecase.VerifyOTPAndRegisterUsecase
}

func NewVerifyOTPAndRegisterController(
	usecase *usecase.VerifyOTPAndRegisterUsecase,
) *VerifyOTPAndRegisterController {
	return &VerifyOTPAndRegisterController{
		usecase: usecase,
	}
}

func RegisterVerifyOTPAndRegisterController(
	r *gin.Engine,
	controller *VerifyOTPAndRegisterController,
) {
	r.POST("/auth/verify-otp", controller.VerifyOTPAndRegister)
}

// VerifyOTPAndRegister godoc
//
//	@Summary		Register user
//	@Description	Verifies OTP, registers the user and returns an access token.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		VerifyOTPAndRegisterControllerInput	true	"Register request"
//	@Success		200		{object}	VerifyOTPAndRegisterControllerOutput
//	@Failure		400		{object}	xerr.SwaggerErrOutput
//	@Failure		500		{object}	xerr.SwaggerErrOutput
//	@Router			/auth/verify-otp [post]
func (c *VerifyOTPAndRegisterController) VerifyOTPAndRegister(ctx *gin.Context) {
	var input VerifyOTPAndRegisterControllerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = ctx.Error(validation.ToValidationError(validationErrs))
			return
		}
		_ = ctx.Error(validation.NewRequestBindingError(
			"verify otp and register",
			validation.JSON,
		))
		return
	}

	output, err := c.usecase.Execute(
		ctx.Request.Context(),
		usecase.VerifyOTPAndRegisterInput{
			Phone:    input.Phone,
			Password: input.Password,
			Code:     input.Code,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, VerifyOTPAndRegisterControllerOutput{
		AccessToken: output.AccessToken,
	})
}
