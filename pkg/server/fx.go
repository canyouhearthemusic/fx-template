package server

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(
			New,
			NewConfig,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *Server) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return server.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return server.Stop(ctx)
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("server")
		}),
	)
}
