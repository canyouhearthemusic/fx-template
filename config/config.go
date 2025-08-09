package config

import (
	"github.com/spf13/viper"
)

const defaultConfigPath = ".env"

func New() *viper.Viper {
	cfg := viper.New()

	cfg.SetConfigFile(defaultConfigPath)
	cfg.SetConfigType("env")
	cfg.AutomaticEnv()

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	return cfg
}
