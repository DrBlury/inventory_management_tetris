package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Config *Config
	Client *redis.Client
}

// NewCache creates a new cache
func NewCache(config *Config) *Cache {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Cache{
		Config: config,
		Client: client,
	}
}

func (c *Cache) GetString(ctx context.Context, key string) (string, error) {
	val, err := c.Client.Get(ctx, "foo").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Cache) SetString(ctx context.Context, key, value string) error {
	err := c.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Invalidate a key
func (c *Cache) Invalidate(ctx context.Context, key string) error {
	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Close() error {
	return c.Client.Close()
}
