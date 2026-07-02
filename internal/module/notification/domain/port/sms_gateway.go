package port

import "context"

type SMSGateway interface {
	Send(ctx context.Context, phone string, message string) error
}
