package cache

import (
	"context"
	"time"
)

type Client interface {
	Cache() Cache
	Close() error
}

type Setter interface {
	Set(context.Context, string, interface{}, time.Duration) (string, error)
}

type Getter interface {
	Get(context.Context, string) (string, error)
}

type Pinger interface {
	Ping(context.Context) error
}

type Cache interface {
	Setter
	Getter
	Pinger
	Close() error
}
