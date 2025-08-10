package database

import (
	"prac/pkg/database/postgres"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"database",
		fx.Options(
			postgres.NewDatabaseClient(),
		),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("database")
		}),
	)
}
