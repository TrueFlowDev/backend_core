package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
)

var (
	ErrSMSDeliveryFailed = xerr.New(
		xerr.CodeInternalError,
		xerr.WithMessage("failed to deliver sms"),
	)
)

type SMSGateway interface {
	Send(ctx context.Context, phone string, message string) error
}
