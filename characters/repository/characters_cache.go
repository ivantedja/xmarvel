package repository

import (
	"context"
	"time"

	"github.com/ivantedja/xmarvel/lib"
)

type cache struct {
	client lib.CacheClient
}

func NewCache(client lib.CacheClient) *cache {
	return &cache{
		client: client,
	}
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration)
	if err != nil {
		return err
	}
	return nil
}
