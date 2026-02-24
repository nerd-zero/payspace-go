package payspace

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newDefaultLogger creates a production zap logger at the given level.
func newDefaultLogger(level zapcore.Level) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "level"
	return cfg.Build()
}

// LoggerConfig holds options for creating a customized zap logger for use with
// the PaySpace client.
type LoggerConfig struct {
	Level       zapcore.Level
	Development bool
	OutputPaths []string
}

// NewLogger creates a zap logger with the given configuration.
func NewLogger(cfg LoggerConfig) (*zap.Logger, error) {
	var zapCfg zap.Config
	if cfg.Development {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
	}
	zapCfg.Level = zap.NewAtomicLevelAt(cfg.Level)
	zapCfg.EncoderConfig.TimeKey = "ts"
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if len(cfg.OutputPaths) > 0 {
		zapCfg.OutputPaths = cfg.OutputPaths
	}
	return zapCfg.Build()
}
