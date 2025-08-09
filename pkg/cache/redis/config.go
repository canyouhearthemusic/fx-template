package redis

import "github.com/spf13/viper"

type Config struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Database int    `mapstructure:"REDIS_DATABASE"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

func NewConfig(provider *viper.Viper) *Config {
	cfg := new(Config)

	if err := provider.Unmarshal(cfg); err != nil {
		panic(err)
	}

	return cfg
}
