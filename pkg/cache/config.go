package cache

import (
	"github.com/spf13/viper"
	"prac/pkg/cache/redis"
)

type Config struct {
	Redis *redis.Config `mapstructure:",squash"`
}

func NewConfig(provider *viper.Viper) *Config {
	cfg := new(Config)

	if err := provider.Unmarshal(cfg); err != nil {
		panic(err)
	}

	return cfg
}
