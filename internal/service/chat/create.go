package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
)

func (s *srv) Create(ctx context.Context, info *model.Chat) (int64, error) {
	id, err := s.repo.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
