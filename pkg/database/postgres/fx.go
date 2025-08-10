package postgres

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewDatabaseClient() fx.Option {
	return fx.Module(
		"postgres",
		fx.Provide(
			NewConfig,
			NewClient,
		),
		fx.Invoke(func(lc fx.Lifecycle, client *Client) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return client.Connect(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return client.Close(ctx)
				},
			})
		}),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("postgres")
		}),
	)
}
