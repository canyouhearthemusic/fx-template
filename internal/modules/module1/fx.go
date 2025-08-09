package module1

import (
	"context"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"module1",
		fx.Provide(
			New,
			NewConfig,
		),
		fx.Invoke(func(lc fx.Lifecycle, service *Service) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					service.SayHi()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
	)
}
