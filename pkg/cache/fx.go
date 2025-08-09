package cache

import (
	"context"
	"prac/pkg/cache/redis"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"cache",
		fx.Provide(
			redis.NewConfig,
			redis.New,
		),
		fx.Invoke(func(lc fx.Lifecycle, cache *redis.Cache) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return cache.Connect(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return cache.Close(ctx)
				},
			})
		}),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("cache")
		}),
	)
}
