package chat

import (
	"context"
	"fmt"

	"github.com/drizzleent/chat-server/internal/client/db"
	"github.com/drizzleent/chat-server/internal/model"
	"github.com/drizzleent/chat-server/internal/repository"
	"github.com/google/uuid"
)

const (
	chatServerTable = "chat_server"
	chatsTable      = "chats"
	id              = "id"
	username        = "username"
	msg             = "message"
	chatUUID        = "chatId"
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

	query := fmt.Sprintf("INSERT INTO %s (%s,%s) values ($1, $2) RETURNING id", chatServerTable, username, msg)

	q := db.Query{
		Name:     "chat.repository.Create",
		QueryRow: query,
	}
	args := []interface{}{chat.Username, chat.Msg}

	var id int64

	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert chat_server: %v", err)
	}

	return id, nil
	//return 0, nil
}
func (r *repo) Delete(ctx context.Context, chatId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", chatServerTable, id)

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

func (r *repo) CreateChat(ctx context.Context, chatId uuid.UUID) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1) RETURNING id", chatsTable, chatUUID)
	q := db.Query{
		Name:     "repository.CreateChat",
		QueryRow: query,
	}

	args := []interface{}{chatId}
	var id int64
	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetChat(ctx context.Context, id string) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE chatId=$1", chatsTable)
	q := db.Query{
		Name:     "repository.GetChat",
		QueryRow: query,
	}
	args := []interface{}{id}
	res, err := r.db.DB().QuaryContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to get chat %v, tag= %v", err, res)
	}

	return nil
}
