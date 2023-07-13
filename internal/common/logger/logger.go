package logger

import (
	"github.com/google/wire"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/pkg/logger"
)

var Set = wire.NewSet(
	NewLoggerAplication,
)

// NewHandler Constructor
func NewLoggerAplication(cfg *config.Configuration) logger.Logger {
	return logger.NewApiLogger(cfg)
}
