package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/zap"
)

type Server struct {
	*fiber.App

	logger *zap.Logger
	cfg    *Config
}

func New(logger *zap.Logger, cfg *Config) *Server {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	}))

	app.Use(recover.New(recover.ConfigDefault))

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	return &Server{
		App:    app,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Server) Start(_ context.Context) error {
	s.logger.Info("Starting server", zap.String("address", s.Addr()))

	go func() {
		if err := s.Listen(s.Addr()); err != nil {
			s.logger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	s.logger.Info("Server is running", zap.String("address", s.Addr()))

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down server...")

	return s.ShutdownWithContext(ctx)
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
}
