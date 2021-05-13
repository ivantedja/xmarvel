package lib

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	client *redis.Client
}

func NewRedis(host, port string) *redisClient {
	r := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return &redisClient{client: r}
}

func (rc *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rc *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := rc.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
