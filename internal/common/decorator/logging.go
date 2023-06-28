package decorator

import (
	"context"
	zapLogger "go-clean-architecture-example/pkg/logger"
	"go.uber.org/zap"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger zapLogger.Logger
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {

	logger := d.logger.WithFiled(zap.String("command", generateActionName(cmd)))
	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.Error("Failed to execute command: " + err.Error())
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger zapLogger.Logger
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := d.logger.WithFiled(zap.String("query", generateActionName(cmd)))
	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Error("Failed to execute query: " + err.Error())
		}
	}()

	return d.base.Handle(ctx, cmd)
}
