package logger

import (
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) (*ZapLogger, error) {
	var zapConfig zap.Config

	if cfg.App.Mode == "dev" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Encoding = "console"

		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
		zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	if err := zapConfig.Level.UnmarshalText([]byte(cfg.Logger.Level)); err != nil {
		return nil, err
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger.Sugar(),
	}, nil
}

func (z *ZapLogger) With(args ...any) port.Logger {
	return &ZapLogger{
		logger: z.logger.With(args...),
	}
}

func (z *ZapLogger) Debug(msg string, args ...any) {
	z.logger.Debugw(msg, args...)
}

func (z *ZapLogger) Info(msg string, args ...any) {
	z.logger.Infow(msg, args...)
}

func (z *ZapLogger) Warn(msg string, args ...any) {
	z.logger.Warnw(msg, args...)
}

func (z *ZapLogger) Error(msg string, args ...any) {
	z.logger.Errorw(msg, args...)
}
