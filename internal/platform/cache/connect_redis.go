package cache

import (
	"context"
	"time"

	"github.com/avast/retry-go/v5"
	"go.uber.org/fx"
)

func ConnectRedis(
	lc fx.Lifecycle,
	redis *Redis,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return retry.New(
				retry.Context(ctx),
				retry.Attempts(10),
				retry.MaxDelay(30*time.Second),
				retry.Delay(time.Second),
				retry.LastErrorOnly(true),
			).Do(func() error {
				return redis.Connect(ctx)
			})
		},
		OnStop: func(ctx context.Context) error {
			return redis.Close()
		},
	})
}
