package adapter

import (
	"context"
	"fmt"
)

type LocalSmsGateway struct{}

func NewLocalSmsGateway() *LocalSmsGateway { return &LocalSmsGateway{} }

func (g *LocalSmsGateway) Send(ctx context.Context, phone string, message string) error {
	fmt.Printf("Sent %q to %q", message, phone)
	return nil
}
