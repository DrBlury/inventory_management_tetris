package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache struct {
	Config *Config
	Client *redis.Client
	log    *zap.SugaredLogger
}

// NewCache creates a new cache
func NewCache(config *Config, log *zap.SugaredLogger) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error("error connecting to cache", zap.Error(err), zap.String("result", result))
		return nil
	}

	log.Info("connected to cache", zap.String("result", result))

	return &Cache{
		Config: config,
		Client: client,
		log:    log,
	}
}

func (c *Cache) GetString(ctx context.Context, key string) (string, error) {
	if c.Client == nil {
		c.log.Error("cache client is nil")
	}
	if key == "" {
		c.log.Error("key is empty")
	}
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		switch err {
		case redis.Nil:
			return "", nil
		default:
			c.log.Error("error getting string from cache", zap.Error(err))
			return "", err
		}
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
