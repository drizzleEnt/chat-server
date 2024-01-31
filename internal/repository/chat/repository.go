package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/client/db"
	"github.com/drizzleent/chat-server/internal/model"
	"github.com/drizzleent/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(context.Context, *model.Chat) (int64, error) {
	return 0, nil
}
func (r *repo) Delete(context.Context, int64) error {
	return nil
}
func (r *repo) Send(context.Context, *model.Chat) error {
	return nil
}
