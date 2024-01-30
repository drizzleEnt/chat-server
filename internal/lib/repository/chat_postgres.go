package repository

import (
	"context"
	"fmt"

	"github.com/drizzleent/chat-server/internal/model"
	"github.com/jackc/pgx/v4"
)

type ChatPostgres struct {
	db *pgx.Conn
}

func NewChatPostgres(db *pgx.Conn) *ChatPostgres {
	return &ChatPostgres{
		db: db,
	}
}

func (r *ChatPostgres) SendMessage(ctx context.Context, user model.User) error {

	var id int

	quary := fmt.Sprintf("INSERT INTO %s (username, message) VALUES ($1, $2) RETURNING id", chatTable)

	row := r.db.QueryRow(ctx, quary, user.Username, user.Message.Text)

	if err := row.Scan(&id); err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}

func (r *ChatPostgres) GetMessage(ctx context.Context) (string, error) {
	return "", nil
}
