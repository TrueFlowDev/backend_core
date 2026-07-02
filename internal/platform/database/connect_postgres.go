package database

import (
	"context"
	"time"

	"github.com/avast/retry-go/v5"
	"go.uber.org/fx"
)

func ConnectPostgres(lc fx.Lifecycle, pg *Postgres) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return retry.New(
				retry.Context(ctx),
				retry.Attempts(10),
				retry.MaxDelay(30*time.Second),
				retry.Delay(time.Second),
				retry.LastErrorOnly(true),
			).Do(
				func() error {
					return pg.Connect()
				},
			)
		},
		OnStop: func(context.Context) error {
			return pg.Close()
		},
	})
}
