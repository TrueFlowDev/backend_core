package middleware

import (
	"errors"
	"net/http"

	"github.com/Ali127Dev/xerr"
	"github.com/gin-gonic/gin"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if xe, ok := errors.AsType[*xerr.Error](err); ok {
			c.JSON(xe.HTTPStatus(), xe)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			xerr.New(
				xerr.CodeInternalError,
				xerr.WithMessage("unknown error"),
			),
		)
	}
}
