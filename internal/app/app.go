package app

import (
	"prac/config"
	"prac/internal/modules/module1"
	"prac/pkg/cache"
	"prac/pkg/database"
	"prac/pkg/logger"
	"prac/pkg/server"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			logger.New,
			validator.New,
		),
		fx.Options(
			cache.NewModule(),
			database.NewModule(),
			server.NewModule(),
			module1.NewModule(),
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
}
