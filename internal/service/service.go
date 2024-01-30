package service

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
)

type ChatService interface {
	Create(context.Context, *model.Chat) (int64, error)
	Delete(context.Context, int64) error
	SendMsg(context.Context, *model.Chat) error
}
