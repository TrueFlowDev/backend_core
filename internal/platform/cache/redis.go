package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client

	host     string
	port     int
	password string
	db       int
}

func NewRedis(cfg *config.Config) *Redis {
	return &Redis{
		host:     cfg.Cache.RedisHost,
		port:     cfg.Cache.RedisPort,
		password: cfg.Cache.RedisPassword,
		db:       cfg.Cache.RedisDb,
	}
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
