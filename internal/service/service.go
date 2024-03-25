package service

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
	"github.com/google/uuid"
)

type ChatService interface {
	GetName(context.Context) (string, error)
	CreateChat(context.Context) (uuid.UUID, error)
	Create(context.Context, *model.Chat) (int64, error)
	Delete(context.Context, int64) error
	GetChat(context.Context, string) error
}
