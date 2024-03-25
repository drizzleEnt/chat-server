package rd

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type rdb struct {
	dbc *redis.Client
}

func NewClient(dbc *redis.Client) *rdb {
	return &rdb{
		dbc: dbc,
	}
}

func (r *rdb) Set(ctx context.Context, key string, value interface{}, exp time.Duration) (string, error) {
	return r.dbc.Set(ctx, key, value, exp).Result()
}

func (r *rdb) Get(ctx context.Context, key string) (string, error) {
	return r.dbc.Get(ctx, key).Result()
}

func (r *rdb) Close() error {
	return r.dbc.Close()
}

func (r *rdb) Ping(ctx context.Context) error {
	return r.dbc.Ping(ctx).Err()
}
