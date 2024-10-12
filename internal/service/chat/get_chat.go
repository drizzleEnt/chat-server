package chat

import (
	"context"
)

func (s *srv) GetChat(ctx context.Context, id string) (bool, error) {
	isExist, err := s.repo.GetChat(ctx, id)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
