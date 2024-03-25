package cache

import (
	"github.com/drizzleent/chat-server/internal/client/cache"
	"github.com/drizzleent/chat-server/internal/repository"
)

type rds struct {
	db cache.Client
}

func NewRepository(db cache.Client) repository.CacheRepository {
	return &rds{
		db: db,
	}
}
