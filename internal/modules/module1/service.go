package module1

import "fmt"

type Service struct {
	cfg *Config
}

func New(cfg *Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) SayHi() {
	fmt.Println("module1 responded:", s.cfg.AppName)
}
