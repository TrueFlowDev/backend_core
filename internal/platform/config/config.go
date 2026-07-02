package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Mode string `env:"MODE" env-default:"dev"`
	Port int    `env:"PORT" env-default:"8080"`
}

type DBConfig struct {
	PostgresHost     string `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" env-required:"true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-required:"true"`
	PostgresDB       string `env:"POSTGRES_DB" env-required:"true"`
}

type CacheConfig struct {
	RedisHost     string `env:"REDIS_HOST" env-default:"localhost"`
	RedisPort     int    `env:"REDIS_PORT" env-default:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" env-required:"true"`
}

type JWTConfig struct {
	AccessSecret string        `env:"ACCESS_SECRET" env-required:"true"`
	AccessTTL    time.Duration `env:"ACCESS_TTL" env-default:"15m"`
}

type OTPConfig struct {
	TTL    time.Duration `env:"TTL" env-default:"2m"`
	Length int           `env:"LENGTH" env-default:"6"`
}

type LoggerConfig struct {
	Level string `env:"LEVEL" env-default:"info"`
}

type Config struct {
	App    AppConfig    `env-prefix:"APP_"`
	DB     DBConfig     `env-prefix:"DB_"`
	Cache  CacheConfig  `env-prefix:"CACHE_"`
	JWT    JWTConfig    `env-prefix:"JWT_"`
	OTP    OTPConfig    `env-prefix:"OTP_"`
	Logger LoggerConfig `env-prefix:"LOGGER_"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
