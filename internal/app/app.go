package app

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"prac/config"
	"prac/internal/modules/module1"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Options(
			module1.NewModule(),
		),
		fx.Provide(
			config.New,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
}
