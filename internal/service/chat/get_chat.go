package chat

import (
	"context"
)

func (s *srv) GetChat(ctx context.Context, id string) error {
	err := s.repo.GetChat(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
