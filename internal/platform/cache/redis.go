package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/avast/retry-go/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Redis struct {
	client *redis.Client

	host     string
	port     int
	password string
	db       int
}

func NewRedis(lc fx.Lifecycle, cfg *config.Config) (*redis.Client, error) {
	rdb := &Redis{
		host:     cfg.Cache.RedisHost,
		port:     cfg.Cache.RedisPort,
		password: cfg.Cache.RedisPassword,
		db:       cfg.Cache.RedisDb,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := retry.New(
		retry.Context(ctx),
		retry.Attempts(10),
		retry.MaxDelay(30*time.Second),
		retry.Delay(time.Second),
		retry.LastErrorOnly(true),
	).Do(func() error {
		return rdb.Connect(ctx)
	})
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})

	return rdb.client, nil
}

func (r *Redis) Connect(ctx context.Context) error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.host, r.port),
		Password: r.password,
		DB:       r.db,

		PoolSize:        20,
		MinIdleConns:    5,
		ConnMaxLifetime: 30 * time.Minute,
		ConnMaxIdleTime: 15 * time.Minute,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return err
	}

	r.client = client

	return nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
