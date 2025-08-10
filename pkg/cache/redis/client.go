package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Client struct {
	*redis.Client

	logger *zap.Logger
	cfg    *Config
}

func NewClient(logger *zap.Logger, cfg *Config) *Client {
	return &Client{
		logger: logger,
		cfg:    cfg,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	c.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port),
		Password: c.cfg.Password,
		DB:       c.cfg.Database,
	})

	c.logger.Info("Connecting to cache...")

	if err := c.Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to cache: %w", err)
	}

	c.logger.Info("Connected to cache")

	return nil
}

func (c *Client) Close(_ context.Context) error {
	c.logger.Info("Closing cache connection...")

	if err := c.Client.Close(); err != nil {
		return fmt.Errorf("failed to close cache: %w", err)
	}

	c.logger.Info("cache connection closed")

	return nil
}

func (c *Client) GetStruct(ctx context.Context, key string, result any) error {
	data, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), result)
}

func (c *Client) SetStruct(ctx context.Context, key string, value any, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(ctx, key, jsonData, ttl).Err()
}
