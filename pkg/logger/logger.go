package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle) (*zap.Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = log.Sync()
			return nil
		},
	})

	return log, nil
}
