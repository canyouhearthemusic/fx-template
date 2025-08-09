package server

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	c := new(Config)

	return c
}
