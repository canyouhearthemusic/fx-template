package postgres

import "github.com/spf13/viper"

type Config struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int    `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Database string `mapstructure:"POSTGRES_DB"`
}

func NewConfig(provider *viper.Viper) *Config {
	cfg := new(Config)

	if err := provider.Unmarshal(cfg); err != nil {
		panic(err)
	}

	return cfg
}
