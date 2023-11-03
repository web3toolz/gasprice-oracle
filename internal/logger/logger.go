package logger

import (
	"fmt"
	"gasprice-oracle/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg config.Config) (*zap.Logger, func()) {

	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	level, err := zapcore.ParseLevel(cfg.LogLevel)

	if err != nil {
		panic(fmt.Sprintf("error while parsing log level %v", err))
	}

	loggerConfig.Level.SetLevel(level)

	logger, err := loggerConfig.Build()

	if err != nil {
		panic(fmt.Sprintf("error while initializing zap logger %v", err))
	}

	return logger, func() {
		_ = logger.Sync()
	}
}
