package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Client struct {
	DB     *pgx.Conn
	cfg    *Config
	logger *zap.Logger
}

func NewClient(logger *zap.Logger, cfg *Config) *Client {
	return &Client{
		cfg:    cfg,
		logger: logger,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.cfg.Host, c.cfg.Port, c.cfg.User, c.cfg.Password, c.cfg.Database)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		c.logger.Fatal("failed to connect to postgres", zap.Error(err))
		return err
	}

	if err := conn.Ping(ctx); err != nil {
		c.logger.Fatal("failed to ping postgres", zap.Error(err))
		return err
	}

	c.DB = conn

	return nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.DB.Close(ctx); err != nil {
		return err
	}

	return nil
}
