package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

const (
	chatTable = "chat_server"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.UserName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = con.Ping(ctx)

	if err != nil {
		return nil, err
	}

	return con, nil
}
