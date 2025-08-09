package redis

type Config struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Database int    `mapstructure:"REDIS_DATABASE"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}
