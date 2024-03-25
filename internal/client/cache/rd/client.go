package rd

import (
	"context"

	"github.com/drizzleent/chat-server/internal/client/cache"
	"github.com/drizzleent/chat-server/internal/config"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	masterDbc cache.Cache
}

func New(ctx context.Context, cfg config.RedisConfig) (cache.Client, error) {

	rClient := redis.NewClient(&redis.Options{
		Addr:       cfg.Address(),
		ClientName: cfg.ClientName(),
		Username:   cfg.UserName(),
		Password:   cfg.Password(),
		DB:         0,
	})

	return &redisClient{
		masterDbc: &rdb{
			dbc: rClient,
		},
	}, nil
}

func (c *redisClient) Cache() cache.Cache {
	return c.masterDbc
}

func (c *redisClient) Close() error {
	if c.masterDbc != nil {
		c.masterDbc.Close()
	}
	return nil
}
