package chat

import (
	"context"
	"fmt"

	"github.com/drizzleent/chat-server/internal/client/db"
	"github.com/drizzleent/chat-server/internal/model"
	"github.com/drizzleent/chat-server/internal/repository"
)

const (
	table    = "chat_server"
	id       = "id"
	username = "username"
	msg      = "message"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {

	query := fmt.Sprintf("INSERT INTO %s (%s,%s) values ($1, $2) RETURNING id", table, username, msg)

	q := db.Query{
		Name:     "chat.repository.Create",
		QueryRow: query,
	}
	args := []interface{}{chat.Name, chat.Msg}

	var id int64

	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert chat_server: %v", err)
	}

	return id, nil
}
func (r *repo) Delete(ctx context.Context, chatId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", table, id)

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRow: query,
	}

	args := []interface{}{chatId}

	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to delete chat %v, tag= %v", err, res)
	}

	return nil
}
func (r *repo) Send(context.Context, *model.Chat) error {
	return nil
}
