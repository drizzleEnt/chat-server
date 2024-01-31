package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Handler func(ctx context.Context) error

type Client interface {
	DB() DB
	Close() error
}

type SQLExecer interface {
	NamedExecer
	QuaryExecer
}

type NamedExecer interface {
	ScanOneContext(context.Context, interface{}, Query, ...interface{}) error
	ScanAllContext(context.Context, interface{}, Query, ...interface{}) error
}

type QuaryExecer interface {
	ExecContext(context.Context, Query, ...interface{}) (pgconn.CommandTag, error)
	QuaryContext(context.Context, Query, ...interface{}) (pgx.Rows, error)
	QuaryRowContext(context.Context, Query, ...interface{}) pgx.Row
}

type Transactor interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type Pinger interface {
	Ping(context.Context) error
}

type Query struct {
	Name     string
	QueryRow string
}

type DB interface {
	SQLExecer
	Pinger
	Transactor
	Close()
}

type TxManager interface {
	ReadCommitted(context.Context, Handler) error
}
