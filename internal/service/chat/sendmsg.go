package chat

import (
	"context"

	"github.com/drizzleent/chat-server/internal/model"
)

func (s *srv) SendMsg(ctx context.Context, info *model.Chat) error {
	err := s.repo.Send(ctx, info)

	if err != nil {
		return err
	}

	return nil
}
