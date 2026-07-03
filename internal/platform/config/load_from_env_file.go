package config

import (
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"github.com/joho/godotenv"
)

func LoadFromEnvFile(logger port.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Warn("envs are not been loaded from .env file")
	}
}
