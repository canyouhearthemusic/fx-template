package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache struct {
	*redis.Client

	logger *zap.Logger
	cfg    *Config
}

func New(logger *zap.Logger, cfg *Config) *Cache {
	return &Cache{
		logger: logger,
		cfg:    cfg,
	}
}

func (c *Cache) Connect(ctx context.Context) error {
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

func (c *Cache) Close(_ context.Context) error {
	c.logger.Info("Closing cache connection...")

	if err := c.Client.Close(); err != nil {
		return fmt.Errorf("failed to close cache: %w", err)
	}

	c.logger.Info("cache connection closed")

	return nil
}

func (c *Cache) GetStruct(ctx context.Context, key string, result any) error {
	data, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), result)
}

func (c *Cache) SetStruct(ctx context.Context, key string, value any, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(ctx, key, jsonData, ttl).Err()
}
