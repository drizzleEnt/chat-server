package pg

import (
	"context"

	"github.com/drizzleent/chat-server/internal/client/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgClient struct {
	masterDbc db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect db: %s", err)
	}

	return &pgClient{
		masterDbc: &pg{
			dbc: dbc,
		},
	}, nil

}

func (c *pgClient) DB() db.DB {
	return c.masterDbc
}

func (c *pgClient) Close() error {
	if c.masterDbc != nil {
		c.masterDbc.Close()
	}

	return nil
}
