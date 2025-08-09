package module1

import "github.com/spf13/viper"

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
}

func NewConfig(provider *viper.Viper) *Config {
	cfg := new(Config)

	if err := provider.Unmarshal(cfg); err != nil {
		panic(err)
	}

	return cfg
}
