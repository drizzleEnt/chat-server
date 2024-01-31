package repository

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
)

type ChatRepository interface {
	Create(context.Context, *model.Chat) (int64, error)
	Delete(context.Context, int64) error
	Send(context.Context, *model.Chat) error
}
