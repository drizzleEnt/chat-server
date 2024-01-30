package repository

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
	"github.com/jackc/pgx/v4"
)

type Chatting interface {
	SendMessage(context.Context, model.Chat) error
	GetMessage(context.Context) (string, error)
}

type Repository struct {
	Chatting
}

func NewReository(db *pgx.Conn) *Repository {
	return &Repository{
		//Chatting: NewChatPostgres(db),
	}
}
