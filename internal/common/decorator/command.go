package decorator

import (
	"context"
	"go-clean-architecture-example/pkg/logger"
)

func ApplyCommandDecorators[H any](handler CommandHandler[H], logger logger.Logger, metricsClient MetricsClient) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: commandMetricsDecorator[H]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
