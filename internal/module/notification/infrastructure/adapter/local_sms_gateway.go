package adapter

import (
	"context"
	"fmt"

	"github.com/TrueFlowDev/Backend/internal/module/notification/domain/port"
)

var _ port.SMSGateway = (*LocalSmsGateway)(nil)

type LocalSmsGateway struct{}

func NewLocalSmsGateway() *LocalSmsGateway { return &LocalSmsGateway{} }

func (g *LocalSmsGateway) Send(ctx context.Context, phone string, message string) error {
	fmt.Printf("Sent %q to %q\n", message, phone)
	return nil
}
