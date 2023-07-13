package decorator

import (
	"context"
	"go-clean-architecture-example/pkg/logger"
)

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger logger.Logger, metricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
