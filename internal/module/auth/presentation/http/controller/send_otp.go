package controller

import (
	"errors"
	"net/http"

	"github.com/TrueFlowDev/Backend/internal/module/auth/application/usecase"
	validation2 "github.com/TrueFlowDev/Backend/internal/platform/server/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SendOtpControllerInput struct {
	Phone string `json:"phone" binding:"required"`
} //	@name	SendOtpInput

type SendOtpController struct {
	usecase *usecase.SendOtpUsecase
}

func NewSendOtpController(usecase *usecase.SendOtpUsecase) *SendOtpController {
	return &SendOtpController{usecase: usecase}
}

func RegisterSendOtpController(r *gin.Engine, sendOtpController *SendOtpController) {
	r.POST("/auth/send-otp", sendOtpController.SendOTP)
}

// SendOTP godoc
//
//	@Summary		Send OTP
//	@Description	Sends a one-time password (OTP) to the specified phone number.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	SendOtpControllerInput	true	"Send OTP request"
//	@Success		204		"OTP sent successfully"
//	@Failure		400		{object}	xerr.SwaggerErrOutput
//	@Failure		500		{object}	xerr.SwaggerErrOutput
//	@Router			/auth/send-otp [post]
func (receiver *SendOtpController) SendOTP(c *gin.Context) {
	var input SendOtpControllerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		if validationErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			_ = c.Error(validation2.ToValidationError(validationErrs))
			return
		}
		_ = c.Error(validation2.NewRequestBindingError("send otp", validation2.JSON))
		return
	}

	err := receiver.usecase.Execute(c.Request.Context(), usecase.SendOtpInput{Phone: input.Phone})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
