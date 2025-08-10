package redis

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewCacheClient() fx.Option {
	return fx.Module(
		"redis",
		fx.Provide(
			NewClient,
			NewConfig,
		),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("redis")
		}),
	)
}
