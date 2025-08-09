package app

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"prac/config"
	"prac/internal/modules/module1"
	"prac/pkg/cache"
	"prac/pkg/logger"
	"prac/pkg/server"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Options(
			server.NewModule(),
			cache.NewModule(),
			module1.NewModule(),
		),
		fx.Provide(
			config.New,
			logger.New,
			validator.New,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
}
