package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/auth"
)

func (s *srv) GetName(ctx context.Context) (string, error) {
	res, err := auth.GetName(ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}
