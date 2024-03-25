package repository

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
	"github.com/google/uuid"
)

type ChatRepository interface {
	Create(context.Context, *model.Chat) (int64, error)
	Delete(context.Context, int64) error
	//Send(context.Context, *model.Chat) error
	CreateChat(context.Context, uuid.UUID) (int64, error)
	GetChat(context.Context, string) error
}

type CacheRepository interface {
}
